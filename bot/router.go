package bot

import (
	"log"

	"github.com/stalkerxxl/telegnom/types"
)

// Router stores handlers, compiles them, and dispatches incoming updates.
type Router struct {
	tmpHandlers []*Handler
	handlers    map[types.UpdateType][]*Handler
}

// newRouter creates a new Router instance.
func newRouter() *Router {
	return &Router{
		handlers: make(map[types.UpdateType][]*Handler),
	}
}

// addHandler stores a handler in a temporary list for later compilation.
func (r *Router) addHandler(h *Handler) {
	r.tmpHandlers = append(r.tmpHandlers, h)
}

// compileHandlers distributes handlers by update type and prepares them for execution.
func (r *Router) compileHandlers() {
	for _, h := range r.tmpHandlers {
		if len(h.typeFilters) == 0 {
			log.Printf("Warning: handler registered without any On*** method. It will be ignored and won't process any updates.")
			continue
		}
		h.compile()
		for t := range h.typeFilters {
			r.handlers[t] = append(r.handlers[t], h)
		}
	}
	r.tmpHandlers = nil
}

// dispatchHandler is called from the global middleware chain to find and execute a suitable handler.
func (r *Router) dispatchHandler(c *Context) {
	updateType := c.Update.Type()

	typeHandlers := r.handlers[updateType]

	// Check handlers for the specific update type
	for _, h := range typeHandlers {
		if runHandler(c, h) {
			return
		}
	}
}

// compileChain assembles middlewares and the final handler into a single function.
func compileChain(middlewares []Middleware, final HandlerFunc) HandlerFunc {
	if final == nil {
		final = func(*Context) {}
	}
	res := final
	for i := len(middlewares) - 1; i >= 0; i-- {
		if middlewares[i] == nil {
			continue
		}
		res = middlewares[i](res)
	}
	return res
}
