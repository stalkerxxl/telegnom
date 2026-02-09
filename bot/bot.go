// Package bot provides functionality for creating and managing Telegram bots.
// It includes structures for handling updates, routing commands, and middleware,
// as well as methods for interacting with the Telegram API via HTTP requests.
// The package supports polling and webhook modes, with options for configuring middleware,
// error handlers, and debugging. Key components: Bot, Router, Context, and related types.
package bot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/stalkerxxl/telegnom/types"
)

const (
	defaultAPIURL         = "https://api.telegram.org"
	defaultPollTimeout    = time.Minute
	defaultUpdatesChanCap = 1024
	defaultWorkers        = 1
	defaultPollLimit      = 100
)

type HandlerFunc func(*Context)
type Filter func(*Context) bool
type Middleware func(HandlerFunc) HandlerFunc
type ErrorHandler func(error)

// Bot represents the main structure of the telegram bot.
type Bot struct {
	apiURL         string
	token          string
	baseURL        string // Pre-computed base URL for requests
	testEnv        bool
	workers        int
	pollTimeout    time.Duration
	clientTimeout  time.Duration
	pollLimit      int
	updatesChanCap int // Capacity of the updates channel
	allowedUpdates types.AllowedUpdates
	client         *http.Client
	debugHandler   HandlerFunc
	errorHandler   ErrorHandler
	middlewares    []Middleware
	router         *Router
	mu             sync.RWMutex
	compileOnce    sync.Once
	// chain is the compiled middleware + handler chain
	chain              HandlerFunc
	updates            chan *types.Update
	lastUpdateID       int64
	webhookSecretToken string
	webhookPath        string
	ctx                context.Context
	producersWG        sync.WaitGroup
}

// NewBot creates a new Bot instance with the provided context and token.
func NewBot(ctx context.Context, token string, opts ...Option) (*Bot, error) {
	if token == "" {
		return nil, errors.New("token required")
	}
	// Basic token validation: Telegram tokens always contain a colon
	if !strings.Contains(token, ":") {
		return nil, errors.New("invalid token format: missing colon")
	}
	if ctx == nil {
		return nil, errors.New("context required")
	}
	b := &Bot{
		apiURL:         defaultAPIURL,
		token:          token,
		workers:        defaultWorkers,
		pollTimeout:    defaultPollTimeout,
		pollLimit:      defaultPollLimit,
		updatesChanCap: defaultUpdatesChanCap,
		debugHandler:   debugHandler,
		errorHandler:   defaultErrorHandler,
		lastUpdateID:   -1,
		ctx:            ctx,
	}
	for _, o := range opts {
		if err := o(b); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	// Smart timeout: if clientTimeout wasn't explicitly set via options,
	// set it to pollTimeout + 30s to avoid premature request cancellation.
	if b.clientTimeout == 0 {
		b.clientTimeout = b.pollTimeout + 30*time.Second
	}

	if b.clientTimeout <= b.pollTimeout {
		return nil, errors.New("client timeout must be greater than poll timeout")
	}

	// Pre-compute baseURL to avoid repeated formatting in every request
	b.baseURL = fmt.Sprintf("%s/bot%s", b.apiURL, b.token)

	// Create a default HTTP client only if one wasn't provided via options
	if b.client == nil {
		b.client = &http.Client{Timeout: b.clientTimeout}
	}

	b.router = newRouter()
	b.updates = make(chan *types.Update, b.updatesChanCap)
	return b, nil
}

// Use adds global middlewares to the bot.
func (b *Bot) Use(mds ...Middleware) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.middlewares = append(b.middlewares, mds...)
}

// Handler registers a new handler through the router.
func (b *Bot) Handler(h HandlerFunc) *Handler {
	hd := &Handler{handlerFunc: h}
	b.router.addHandler(hd)
	return hd
}

// Group creates a new handler group.
func (b *Bot) Group() *Group {
	return &Group{bot: b}
}

// compile performs handler compilation and assembles the global middleware chain.
func (b *Bot) compile() {
	b.compileOnce.Do(func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		b.router.compileHandlers()
		b.chain = compileChain(b.middlewares, b.router.dispatchHandler)
	})
}

// error is a centralized error reporting method.
func (b *Bot) error(err error) {
	if b.errorHandler != nil {
		b.errorHandler(err)
	}
}

// defaultErrorHandler provides basic error logging.
func defaultErrorHandler(err error) {
	log.Printf("ERROR: %v", err)
}

// debugHandler provides basic update logging in JSON format.
func debugHandler(tg *Context) {
	data, err := json.Marshal(tg.Update)
	if err != nil {
		log.Printf("DEBUG ERROR: %v", err)
		return
	}
	log.Printf("UPDATE:%s", string(data))
}

// panicRecovery is a middleware that recovers from panics in handlers.
func panicRecovery(next HandlerFunc) HandlerFunc {
	return func(tg *Context) {
		defer func() {
			if r := recover(); r != nil {
				tg.Bot.error(fmt.Errorf("panic recovered: %v", r))
			}
		}()
		next(tg)
	}
}
