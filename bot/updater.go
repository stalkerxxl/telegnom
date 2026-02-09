package bot

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/stalkerxxl/telegnom/types"
)

const maxBackoffDelay = time.Second * 5

// getUpdatesParams represents the parameters for the getUpdates API method.
type getUpdatesParams struct {
	Offset         int64                `json:"offset,omitempty"`
	Limit          int                  `json:"limit,omitempty"`
	Timeout        int                  `json:"timeout,omitempty"`
	AllowedUpdates types.AllowedUpdates `json:"allowed_updates,omitempty"`
}

// StartPolling starts the long-polling loop and the specified number of worker goroutines.
// This method is blocking and will run until the context provided to NewBot is canceled.
func (b *Bot) StartPolling() {
	// Register polling as a producer
	b.producersWG.Add(1)
	go b.getUpdates()

	// Run will start workers and block until context is canceled
	_ = b.Run()
}

// StartWebhook starts an HTTP server and the specified number of worker
// goroutines. This method is blocking and will run until the context provided to
// [NewBot] is canceled or the server returns an error.
func (b *Bot) StartWebhook(addr string) error {
	return b.startWebhook(addr, "", "")
}

// StartWebhookTLS starts an HTTPS server and the specified number of worker goroutines.
func (b *Bot) StartWebhookTLS(addr, certFile, keyFile string) error {
	return b.startWebhook(addr, certFile, keyFile)
}

func (b *Bot) startWebhook(addr, certFile, keyFile string) error {
	mux := http.NewServeMux()

	path := "/"
	if b.webhookPath != "" {
		path = b.webhookPath
	}

	mux.HandleFunc(path, b.WebhookHandler())

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return b.ctx
		},
	}

	// Register server as a producer
	b.producersWG.Add(1)

	// Channel to capture server error
	errChan := make(chan error, 1)

	// Start server in a separate goroutine
	go func() {
		var err error
		if certFile != "" && keyFile != "" {
			err = srv.ListenAndServeTLS(certFile, keyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	// Graceful shutdown monitor
	go func() {
		defer b.producersWG.Done()
		<-b.ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			b.error(fmt.Errorf("webhook shutdown error: %w", err))
		}
	}()

	// Start bot runners
	botErr := b.Run()

	// Return server error if any, otherwise bot error
	select {
	case err := <-errChan:
		return err
	default:
		return botErr
	}
}

// Run starts the worker goroutines to process updates and blocks until the
// context is canceled. Use this method when you are receiving updates from an
// external source (e.g. your own web server).
//
// Graceful shutdown note: If you are using an external HTTP server, you MUST stop the server
// (ensure no new requests are being accepted) BEFORE canceling the bot's context.
// This ensures that all currently processing updates are delivered to the workers before
// the updates channel is closed.
func (b *Bot) Run() error {
	b.compile()

	var wg sync.WaitGroup
	// Start worker goroutines to process updates concurrently
	wg.Add(b.workers)
	for i := 0; i < b.workers; i++ {
		go b.waitUpdates(&wg)
	}

	// Wait for context cancellation
	<-b.ctx.Done()

	// Wait for all producers (Polling, Webhook) to stop sending updates
	b.producersWG.Wait()

	// Signal workers to stop by closing the channel
	close(b.updates)
	wg.Wait()

	return b.ctx.Err()
}

// WebhookHandler returns a [net/http.HandlerFunc] to be used with a web server.
// It handles incoming updates from Telegram via Webhook.
func (b *Bot) WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Register active request in producersWG to prevent channel closure
		b.producersWG.Add(1)
		defer b.producersWG.Done()

		// Quick check: if bot is stopping, reject request immediately
		select {
		case <-b.ctx.Done():
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		default:
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if b.webhookSecretToken != "" {
			token := r.Header.Get("X-Telegram-Bot-Api-Secret-Token")
			if subtle.ConstantTimeCompare([]byte(token), []byte(b.webhookSecretToken)) != 1 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		// Limit body size to 10MB to prevent DoS attacks
		r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

		var update types.Update
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			b.error(fmt.Errorf("webhook decode error: %w", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		select {
		case b.updates <- &update:
			w.WriteHeader(http.StatusOK)
		case <-b.ctx.Done():
			w.WriteHeader(http.StatusServiceUnavailable)
		default:
			// Channel is full
			b.error(fmt.Errorf("webhook: updates channel full, update %d dropped", update.ID))
			w.WriteHeader(http.StatusTooManyRequests)
		}
	}
}

// getUpdates fetches new updates from the Telegram API in a loop.
// It handles API errors with an exponential backoff mechanism.
func (b *Bot) getUpdates() {
	defer b.producersWG.Done()
	var retryDelay time.Duration

	for {
		// Check if context is already canceled before making a request
		select {
		case <-b.ctx.Done():
			return
		default:
		}

		// Wait before retrying if a previous request failed
		if retryDelay > 0 {
			select {
			case <-b.ctx.Done():
				return
			case <-time.After(retryDelay):
			}
		}

		p := &getUpdatesParams{
			Timeout: int((b.pollTimeout - time.Second).Seconds()),
			Offset:  atomic.LoadInt64(&b.lastUpdateID) + 1,
			Limit:   b.pollLimit,
		}
		if b.allowedUpdates != nil {
			p.AllowedUpdates = b.allowedUpdates
		}

		var updates []*types.Update
		err := b.apiRequest("getUpdates", p, &updates)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}
			b.error(fmt.Errorf("getUpdates error: %w", err))
			retryDelay = nextDelay(retryDelay)
			continue
		}
		retryDelay = 0

		// Forward received updates to the workers via the channel
		for _, upd := range updates {
			atomic.StoreInt64(&b.lastUpdateID, upd.ID)
			select {
			case <-b.ctx.Done():
				return
			case b.updates <- upd:
			}
		}
	}
}

// waitUpdates listens for updates from the internal channel and passes them to processUpdate.
// It terminates automatically when the updates channel is closed by getUpdates.
func (b *Bot) waitUpdates(wg *sync.WaitGroup) {
	defer wg.Done()
	for upd := range b.updates {
		b.processUpdate(upd)
	}
}

// processUpdate creates a new Context for the update and executes the middleware chain.
// It includes a recovery mechanism to prevent worker termination in case of a crash in user handlers.
func (b *Bot) processUpdate(upd *types.Update) {
	c := newContext(b, upd)
	if b.debugHandler != nil {
		b.debugHandler(c)
	}
	if b.chain != nil {
		b.chain(c)
	}
}

// nextDelay computes the duration to wait before the next API attempt using an exponential backoff.
func nextDelay(prev time.Duration) time.Duration {
	switch {
	case prev <= 0:
		return 100 * time.Millisecond
	case prev >= maxBackoffDelay:
		return maxBackoffDelay
	default:
		next := prev * 2
		if next > maxBackoffDelay {
			return maxBackoffDelay
		}
		return next
	}
}
