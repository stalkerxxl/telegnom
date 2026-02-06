package bot

import (
	"context"
	"testing"

	"github.com/stalkerxxl/telegnom/types"
)

func TestGroupMiddlewareAndFilterOrder(t *testing.T) {
	b, _ := NewBot(context.Background(), "123:test")

	executionOrder := ""
	add := func(s string) {
		if executionOrder != "" {
			executionOrder += " -> "
		}
		executionOrder += s
	}

	// Global Middleware
	b.Use(func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			add("GlobalMiddleware_IN")
			next(c)
			add("GlobalMiddleware_OUT")
		}
	})

	// Create Group
	adminGroup := b.Group()

	// Group Middleware
	adminGroup.Use(func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			add("GroupMiddleware_IN")
			next(c)
			add("GroupMiddleware_OUT")
		}
	})

	// Handler in Group
	adminGroup.Handler(func(c *Context) {
		add("Handler")
	}).Use(func(next HandlerFunc) HandlerFunc {
		// Local Middleware
		return func(c *Context) {
			add("HandlerMiddleware_IN")
			next(c)
			add("HandlerMiddleware_OUT")
		}
	}).OnMessage(func(c *Context) bool {
		// Local Filter
		add("HandlerFilter")
		return true
	})

	// Compile and run
	b.compile()
	b.mu.Lock()
	b.chain = compileChain(b.middlewares, b.router.dispatchHandler)
	b.mu.Unlock()
	ctx := newContext(b, &types.Update{Message: &types.Message{Text: "test"}})

	if ch := b.chain; ch != nil {
		ch(ctx)
	}

	expected := "GlobalMiddleware_IN -> HandlerFilter -> " +
		"GroupMiddleware_IN -> HandlerMiddleware_IN -> " +
		"Handler -> HandlerMiddleware_OUT -> " +
		"GroupMiddleware_OUT -> GlobalMiddleware_OUT"

	if executionOrder != expected {
		t.Errorf("\nWrong execution order!\nExpected: %s\nGot: %s", expected, executionOrder)
	}
}

func TestEmptyGroup(t *testing.T) {
	b, _ := NewBot(context.Background(), "123:test")

	executed := false

	// Create empty group
	emptyGroup := b.Group()

	// Handler without group settings
	emptyGroup.Handler(func(c *Context) {
		executed = true
	}).OnMessage()

	// Compile and run
	b.compile()
	b.mu.Lock()
	b.chain = compileChain(b.middlewares, b.router.dispatchHandler)
	b.mu.Unlock()
	ctx := newContext(b, &types.Update{Message: &types.Message{Text: "test"}})

	if ch := b.chain; ch != nil {
		ch(ctx)
	}

	if !executed {
		t.Error("Handler in empty group should execute")
	}
}

func TestGroupNilMiddleware(t *testing.T) {
	b, _ := NewBot(context.Background(), "123:test")

	executed := false

	group := b.Group()
	// Add nil middleware
	group.Use(nil, func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			executed = true
			next(c)
		}
	}, nil)

	group.Handler(func(c *Context) {
	}).OnMessage()

	// Compile and run
	b.compile()
	b.mu.Lock()
	b.chain = compileChain(b.middlewares, b.router.dispatchHandler)
	b.mu.Unlock()
	ctx := newContext(b, &types.Update{Message: &types.Message{Text: "test"}})

	if ch := b.chain; ch != nil {
		ch(ctx)
	}

	if !executed {
		t.Error("Non-nil middleware should execute")
	}
}

func TestMultipleGroupSettings(t *testing.T) {
	b, _ := NewBot(context.Background(), "123:test")

	order := ""

	group := b.Group()
	group.Use(func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			order += "M1"
			next(c)
		}
	}, func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			order += "M2"
			next(c)
		}
	})

	group.Handler(func(c *Context) {
		order += "H"
	}).OnMessage()

	// Compile and run
	b.compile()
	b.mu.Lock()
	b.chain = compileChain(b.middlewares, b.router.dispatchHandler)
	b.mu.Unlock()
	ctx := newContext(b, &types.Update{Message: &types.Message{Text: "test"}})

	if ch := b.chain; ch != nil {
		ch(ctx)
	}

	expected := "M1M2H"
	if order != expected {
		t.Errorf("Wrong order: expected %s, got %s", expected, order)
	}
}

func TestGroupIsolation(t *testing.T) {
	b, _ := NewBot(context.Background(), "123:test")

	order := ""

	group1 := b.Group()
	group1.Use(func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			order += "G1"
			next(c)
		}
	})

	group2 := b.Group()
	group2.Use(func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			order += "G2"
			next(c)
		}
	})

	group1.Handler(func(c *Context) {
		order += "H1"
	}).OnMessage()

	group2.Handler(func(c *Context) {
		order += "H2"
	}).OnMessage()

	// Compile and run
	b.compile()
	b.mu.Lock()
	b.chain = compileChain(b.middlewares, b.router.dispatchHandler)
	b.mu.Unlock()
	ctx := newContext(b, &types.Update{Message: &types.Message{Text: "test"}})

	if ch := b.chain; ch != nil {
		ch(ctx)
	}

	// Since both handlers match, but router picks first, expect G1H1 or G2H2
	if order != "G1H1" && order != "G2H2" {
		t.Errorf("Unexpected order: %s", order)
	}
}

func TestGroupHandlerChaining(t *testing.T) {
	b, _ := NewBot(context.Background(), "123:test")

	order := ""

	group := b.Group()
	group.Use(func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			order += "GM"
			next(c)
		}
	})

	handler := group.Handler(func(c *Context) {
		order += "H"
	})
	handler.Use(func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			order += "LM"
			next(c)
		}
	}).OnMessage()

	// Compile and run
	b.compile()
	b.mu.Lock()
	b.chain = compileChain(b.middlewares, b.router.dispatchHandler)
	b.mu.Unlock()
	ctx := newContext(b, &types.Update{Message: &types.Message{Text: "test"}})

	if ch := b.chain; ch != nil {
		ch(ctx)
	}

	expected := "GMLMH" // Group MW, Local MW, Handler
	if order != expected {
		t.Errorf("Wrong chaining order: expected %s, got %s", expected, order)
	}
}
