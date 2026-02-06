package bot

// Group represents a collection of handlers that share common middlewares.
type Group struct {
	bot         *Bot
	middlewares []Middleware
}

// Use adds group-level middlewares that will be applied to all handlers in this group.
func (gr *Group) Use(m ...Middleware) *Group {
	for _, mw := range m {
		if mw != nil {
			gr.middlewares = append(gr.middlewares, mw)
		}
	}
	return gr
}

// Handler registers a new handler within the group.
// The handler will inherit all middlewares defined in this group.
func (gr *Group) Handler(h HandlerFunc) *Handler {
	hd := gr.bot.Handler(h)
	hd.group = gr
	return hd
}
