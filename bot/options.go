package bot

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stalkerxxl/telegnom/types"
)

// Option is a function that modifies the Bot configuration.
type Option func(b *Bot) error

// WithAPIURL sets an alternative API URL.
func WithAPIURL(apiUrl string) Option {
	return func(b *Bot) error {
		if _, err := url.ParseRequestURI(apiUrl); err != nil {
			return fmt.Errorf("invalid api url: %w", err)
		}
		b.apiURL = apiUrl
		return nil
	}
}

// WithNumWorkers sets the number of worker goroutines.
func WithNumWorkers(workers int) Option {
	return func(b *Bot) error {
		if workers <= 0 {
			return fmt.Errorf("num workers must be greater than 0; got %d", workers)
		}
		b.workers = workers
		return nil
	}
}

// WithAllowedUpdates sets the list of allowed update types.
func WithAllowedUpdates(ut ...types.UpdateType) Option {
	return func(b *Bot) error {
		if len(ut) == 0 {
			return nil
		}
		// ensure uniqueness
		seen := make(map[types.UpdateType]struct{})
		for _, t := range b.allowedUpdates {
			seen[t] = struct{}{}
		}
		for _, t := range ut {
			if _, ok := seen[t]; !ok {
				b.allowedUpdates = append(b.allowedUpdates, t)
				seen[t] = struct{}{}
			}
		}
		return nil
	}
}

// WithPollTimeout sets the long-polling timeout.
func WithPollTimeout(d time.Duration) Option {
	return func(b *Bot) error {
		if d <= 0 {
			return fmt.Errorf("poll timeout must be greater than 0; got %v", d)
		}
		b.pollTimeout = d
		return nil
	}
}

// WithPollLimit sets the maximum number of updates to be retrieved per request.
//
//goland:noinspection GoUnusedExportedFunction
func WithPollLimit(limit int) Option {
	return func(b *Bot) error {
		if limit < 1 || limit > 100 {
			return fmt.Errorf("poll limit must be between 1 and 100; got %d", limit)
		}
		b.pollLimit = limit
		return nil
	}
}

// WithPollOffset sets the identifier of the first update to be returned. This
// corresponds to the "offset" parameter in the TelegramBot API.
//
// Internal logic: The library sets the internal counter to (offset - 1)
// because the main polling loop always adds 1 to the last received update ID.
// This ensures that the exact value passed to this option is used in the first API request.
//
// Use cases:
//   - Positive number (e.g., 814781303): The bot starts retrieving updates starting from this specific ID.
//   - Negative number (e.g., -1): The bot retrieves only the most recent update in the queue.
//     All previous updates will be automatically marked as confirmed by Telegram.
//
// Warning: Telegram stores updates for only 24 hours. Once you call getUpdates with offset N+1,
// all updates with IDs less than or equal to N are permanently deleted from the Telegram servers.
//
//goland:noinspection GoUnusedExportedFunction
func WithPollOffset(offset int64) Option {
	return func(b *Bot) error {
		b.lastUpdateID = offset - 1
		return nil
	}
}

// WithClientTimeout sets the HTTP client timeout.
// It should be larger than the value in WithPollTimeout.
func WithClientTimeout(d time.Duration) Option {
	return func(b *Bot) error {
		if d <= 0 {
			return fmt.Errorf("client timeout must be greater than 0; got %v", d)
		}
		b.clientTimeout = d
		return nil
	}
}

// WithUpdatesChanCap sets the capacity of the updates channel.
//
//goland:noinspection GoUnusedExportedFunction
func WithUpdatesChanCap(cap int) Option {
	return func(b *Bot) error {
		if cap < 1 {
			return fmt.Errorf("updates channel capacity must be greater than 0; got %d", cap)
		}
		b.updatesChanCap = cap
		return nil
	}
}

// WithHTTPClient sets a custom HTTP client.
//
//goland:noinspection GoUnusedExportedFunction
func WithHTTPClient(client *http.Client) Option {
	return func(b *Bot) error {
		if client == nil {
			return fmt.Errorf("http client cannot be nil")
		}
		b.client = client
		return nil
	}
}

// WithDebugHandler sets a custom debug handler.
// To disable debug, pass nil.
func WithDebugHandler(h HandlerFunc) Option {
	return func(b *Bot) error {
		b.debugHandler = h
		return nil
	}
}

// WithErrorHandler sets a custom error handler.
// To disable error handling, pass nil.
func WithErrorHandler(h ErrorHandler) Option {
	return func(b *Bot) error {
		b.errorHandler = h
		return nil
	}
}

// WithTestEnv sets the bot to use the Telegram Bot API test environment.
//
// See https://core.telegram.org/bots/features#dedicated-test-environment
//
//goland:noinspection GoUnusedExportedFunction
func WithTestEnv() Option {
	return func(b *Bot) error {
		b.testEnv = true
		return nil
	}
}

// WithPanicRecovery adds a global middleware that recovers from panics.
func WithPanicRecovery() Option {
	return func(b *Bot) error {
		b.Use(panicRecovery)
		return nil
	}
}

// WithWebhookSecretToken sets the secret token used for verification of webhook requests.
func WithWebhookSecretToken(token string) Option {
	return func(b *Bot) error {
		b.webhookSecretToken = token
		return nil
	}
}

// WithWebhookURL sets the webhook path based on the provided URL.
// The path will be extracted from the URL and used in the internal router.
// Example: https://example.com/bot -> path will be "/bot"
func WithWebhookURL(webhookURL string) Option {
	return func(b *Bot) error {
		u, err := url.Parse(webhookURL)
		if err != nil {
			return fmt.Errorf("invalid webhook url: %w", err)
		}
		path := u.Path
		if path == "" {
			path = "/"
		}
		b.webhookPath = path
		return nil
	}
}
