package bot

import (
	"context"
	"testing"
	"time"

	"github.com/stalkerxxl/telegnom/types"
)

func TestNewBotBasic(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	b, err := NewBot(ctx, "123:test_token")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b == nil {
		t.Fatal("Expected bot to be created")
	}
	if b.token != token {
		t.Errorf("Expected token %s, got %s", token, b.token)
	}
	if b.apiURL != defaultAPIURL {
		t.Errorf("Expected apiURL %s, got %s", defaultAPIURL, b.apiURL)
	}
	if b.workers != defaultWorkers {
		t.Errorf("Expected workers %d, got %d", defaultWorkers, b.workers)
	}
	if b.pollTimeout != defaultPollTimeout {
		t.Errorf("Expected pollTimeout %v, got %v", defaultPollTimeout, b.pollTimeout)
	}
	if b.clientTimeout != defaultPollTimeout+30*time.Second {
		t.Errorf("Expected clientTimeout %v, got %v", defaultPollTimeout+30*time.Second, b.clientTimeout)
	}
	if b.debugHandler == nil {
		t.Error("Expected default debugHandler to be set")
	}
	if b.allowedUpdates != nil {
		t.Errorf("Expected allowedUpdates nil, got %v", b.allowedUpdates)
	}
	if b.client == nil {
		t.Error("Expected client to be set")
	}
	if b.client.Timeout != b.clientTimeout {
		t.Errorf("Expected client timeout %v, got %v", b.clientTimeout, b.client.Timeout)
	}
	if b.router == nil {
		t.Error("Expected router to be set")
	}
	if b.updates == nil {
		t.Error("Expected updates channel to be set")
	}
	if cap(b.updates) != defaultUpdatesChanCap {
		t.Errorf("Expected updates cap %d, got %d", defaultUpdatesChanCap, cap(b.updates))
	}
	if b.lastUpdateID != -1 {
		t.Errorf("Expected lastUpdateID 0, got %d", b.lastUpdateID)
	}
	if b.ctx != ctx {
		t.Error("Expected ctx to be set to provided context")
	}
	if len(b.middlewares) != 0 {
		t.Errorf("Expected middlewares len 0, got %d", len(b.middlewares))
	}
	if b.chain != nil {
		t.Error("Expected chain to be nil")
	}
	if b.debugHandler == nil {
		t.Error("Expected debugHandler to be set")
	}
	if b.errorHandler == nil {
		t.Error("Expected errorHandler to be set")
	}
}

func TestNewBotInvalidToken(t *testing.T) {
	ctx := context.Background()
	_, err := NewBot(ctx, "")
	if err == nil {
		t.Fatal("Expected error for empty token")
	}
	if err.Error() != "token required" {
		t.Errorf("Expected error 'token required', got %v", err)
	}
}

func TestNewBotInvalidContext(t *testing.T) {
	token := "123:test_token"
	_, err := NewBot(nil, token)
	if err == nil {
		t.Fatal("Expected error for nil context")
	}
	if err.Error() != "context required" {
		t.Errorf("Expected error 'context required', got %v", err)
	}
}

func TestNewBotWithApiUrl(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	customUrl := "https://custom.api.telegram.org"
	b, err := NewBot(ctx, token, WithAPIURL(customUrl))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b.apiURL != customUrl {
		t.Errorf("Expected apiURL %s, got %s", customUrl, b.apiURL)
	}
}

func TestNewBotWithNumWorkers(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	workers := 5
	b, err := NewBot(ctx, token, WithNumWorkers(workers))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b.workers != workers {
		t.Errorf("Expected workers %d, got %d", workers, b.workers)
	}
}

func TestNewBotWithPollTimeout(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	pollTimeout := 30 * time.Second
	b, err := NewBot(ctx, token, WithPollTimeout(pollTimeout))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b.pollTimeout != pollTimeout {
		t.Errorf("Expected pollTimeout %v, got %v", pollTimeout, b.pollTimeout)
	}
}

func TestNewBotWithClientTimeout(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	clientTimeout := 3 * time.Minute
	b, err := NewBot(ctx, token, WithClientTimeout(clientTimeout))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b.clientTimeout != clientTimeout {
		t.Errorf("Expected clientTimeout %v, got %v", clientTimeout, b.clientTimeout)
	}
	// Test invalid case: clientTimeout <= pollTimeout
	_, err = NewBot(ctx, token, WithPollTimeout(2*time.Minute), WithClientTimeout(1*time.Minute))
	if err == nil {
		t.Fatal("Expected error when clientTimeout <= pollTimeout")
	}
}

func TestNewBotWithAllowedUpdates(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	allowed := []types.UpdateType{types.MessageType, types.CallbackQueryType}
	b, err := NewBot(ctx, token, WithAllowedUpdates(allowed...))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(b.allowedUpdates) != 2 {
		t.Errorf("Expected 2 allowed updates, got %d", len(b.allowedUpdates))
	}
	if b.allowedUpdates[0] != types.MessageType {
		t.Errorf("Expected first update %s, got %s", types.MessageType, b.allowedUpdates[0])
	}
}

func TestNewBotDebugDisabled(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	// Now we disable debug by passing nil
	b, err := NewBot(ctx, token, WithDebugHandler(nil))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b.debugHandler != nil {
		t.Errorf("Expected debugHandler to be nil, got %v", b.debugHandler)
	}
}

func TestNewBotWithDebugHandler(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	customHandler := func(c *Context) {
		// custom debug handler
	}
	b, err := NewBot(ctx, token, WithDebugHandler(customHandler))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b.debugHandler == nil {
		t.Error("Expected debugHandler to be set")
	}
}

func TestNewBotWithErrorHandler(t *testing.T) {
	ctx := context.Background()
	token := "123:test_token"
	// Test setting a custom handler
	customHandler := func(err error) {
		// custom error handler
	}
	b, err := NewBot(ctx, token, WithErrorHandler(customHandler))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b.errorHandler == nil {
		t.Error("Expected errorHandler to be set")
	}

	// Test disabling the handler
	b2, err := NewBot(ctx, token, WithErrorHandler(nil))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b2.errorHandler != nil {
		t.Error("Expected errorHandler to be nil when disabled")
	}
}
