package bot

import (
	"github.com/stalkerxxl/telegnom/types"
)

// Handler represents a single update handler with its own filters and middleware.
type Handler struct {
	group          *Group                        // Reference to the parent group (if any)
	handlerFunc    HandlerFunc                   // The actual update processing logic
	wrappedHandler HandlerFunc                   // The final middleware-wrapped handler function
	typeFilters    map[types.UpdateType][]Filter // Filters applied to specific update types
	middlewares    []Middleware                  // Local middleware specific to this handler
}

func (h *Handler) addTypeFilter(t types.UpdateType, f ...Filter) {
	if h.typeFilters == nil {
		h.typeFilters = make(map[types.UpdateType][]Filter)
	}
	h.typeFilters[t] = append(h.typeFilters[t], f...)
}

// OnMessage registers the handler for Message updates and adds optional filters.
func (h *Handler) OnMessage(f ...Filter) *Handler {
	h.addTypeFilter(types.MessageType, f...)
	return h
}

// OnEditedMessage registers the handler for EditedMessage updates and adds optional filters.
func (h *Handler) OnEditedMessage(f ...Filter) *Handler {
	h.addTypeFilter(types.EditedMessageType, f...)
	return h
}

// OnChannelPost registers the handler for ChannelPost updates and adds optional filters.
func (h *Handler) OnChannelPost(f ...Filter) *Handler {
	h.addTypeFilter(types.ChannelPostType, f...)
	return h
}

// OnEditedChannelPost registers the handler for EditedChannelPost updates and adds optional filters.
func (h *Handler) OnEditedChannelPost(f ...Filter) *Handler {
	h.addTypeFilter(types.EditedChannelPostType, f...)
	return h
}

// OnBusinessConnection registers the handler for BusinessConnection updates and adds optional filters.
func (h *Handler) OnBusinessConnection(f ...Filter) *Handler {
	h.addTypeFilter(types.BusinessConnectionType, f...)
	return h
}

// OnBusinessMessage registers the handler for BusinessMessage updates and adds optional filters.
func (h *Handler) OnBusinessMessage(f ...Filter) *Handler {
	h.addTypeFilter(types.BusinessMessageType, f...)
	return h
}

// OnEditedBusinessMessage registers the handler for EditedBusinessMessage updates and adds optional filters.
func (h *Handler) OnEditedBusinessMessage(f ...Filter) *Handler {
	h.addTypeFilter(types.EditedBusinessMessageType, f...)
	return h
}

// OnDeletedBusinessMessages registers the handler for DeletedBusinessMessages updates and adds optional filters.
func (h *Handler) OnDeletedBusinessMessages(f ...Filter) *Handler {
	h.addTypeFilter(types.DeletedBusinessMessagesType, f...)
	return h
}

// OnMessageReaction registers the handler for MessageReaction updates and adds optional filters.
func (h *Handler) OnMessageReaction(f ...Filter) *Handler {
	h.addTypeFilter(types.MessageReactionType, f...)
	return h
}

// OnMessageReactionCount registers the handler for MessageReactionCount updates and adds optional filters.
func (h *Handler) OnMessageReactionCount(f ...Filter) *Handler {
	h.addTypeFilter(types.MessageReactionCountType, f...)
	return h
}

// OnInlineQuery registers the handler for InlineQuery updates and adds optional filters.
func (h *Handler) OnInlineQuery(f ...Filter) *Handler {
	h.addTypeFilter(types.InlineQueryType, f...)
	return h
}

// OnChosenInlineResult registers the handler for ChosenInlineResult updates and adds optional filters.
func (h *Handler) OnChosenInlineResult(f ...Filter) *Handler {
	h.addTypeFilter(types.ChosenInlineResultType, f...)
	return h
}

// OnCallbackQuery registers the handler for CallbackQuery updates and adds optional filters.
func (h *Handler) OnCallbackQuery(f ...Filter) *Handler {
	h.addTypeFilter(types.CallbackQueryType, f...)
	return h
}

// OnShippingQuery registers the handler for ShippingQuery updates and adds optional filters.
func (h *Handler) OnShippingQuery(f ...Filter) *Handler {
	h.addTypeFilter(types.ShippingQueryType, f...)
	return h
}

// OnPreCheckoutQuery registers the handler for PreCheckoutQuery updates and adds optional filters.
func (h *Handler) OnPreCheckoutQuery(f ...Filter) *Handler {
	h.addTypeFilter(types.PreCheckoutQueryType, f...)
	return h
}

// OnPurchasedPaidMedia registers the handler for PurchasedPaidMedia updates and adds optional filters.
func (h *Handler) OnPurchasedPaidMedia(f ...Filter) *Handler {
	h.addTypeFilter(types.PurchasedPaidMediaType, f...)
	return h
}

// OnPoll registers the handler for Poll updates and adds optional filters.
func (h *Handler) OnPoll(f ...Filter) *Handler {
	h.addTypeFilter(types.PollType, f...)
	return h
}

// OnPollAnswer registers the handler for PollAnswer updates and adds optional filters.
func (h *Handler) OnPollAnswer(f ...Filter) *Handler {
	h.addTypeFilter(types.PollAnswerType, f...)
	return h
}

// OnMyChatMember registers the handler for MyChatMember updates and adds optional filters.
func (h *Handler) OnMyChatMember(f ...Filter) *Handler {
	h.addTypeFilter(types.MyChatMemberType, f...)
	return h
}

// OnChatMember registers the handler for ChatMember updates and adds optional filters.
func (h *Handler) OnChatMember(f ...Filter) *Handler {
	h.addTypeFilter(types.ChatMemberType, f...)
	return h
}

// OnChatJoinRequest registers the handler for ChatJoinRequest updates and adds optional filters.
func (h *Handler) OnChatJoinRequest(f ...Filter) *Handler {
	h.addTypeFilter(types.ChatJoinRequestType, f...)
	return h
}

// OnChatBoost registers the handler for ChatBoost updates and adds optional filters.
func (h *Handler) OnChatBoost(f ...Filter) *Handler {
	h.addTypeFilter(types.ChatBoostType, f...)
	return h
}

// OnRemovedChatBoost registers the handler for RemovedChatBoost updates and adds optional filters.
func (h *Handler) OnRemovedChatBoost(f ...Filter) *Handler {
	h.addTypeFilter(types.RemovedChatBoostType, f...)
	return h
}

// Use adds local middleware to the handler.
func (h *Handler) Use(m ...Middleware) *Handler {
	h.middlewares = append(h.middlewares, m...)
	return h
}

// compile wraps the main handlerFunc into a "layers" structure:
// 1. Group Middlewares
// 2. Handler-specific Type Filters
// 3. Local Middlewares
// 4. The handler function itself
func (h *Handler) compile() {
	if h.handlerFunc == nil {
		panic("handler: handlerFunc cannot be nil. Make sure you provided a function to the Handler method.")
	}

	// 1. Start with the core handler function
	final := h.handlerFunc

	// 2. Wrap with Local Middlewares
	for i := len(h.middlewares) - 1; i >= 0; i-- {
		if h.middlewares[i] != nil {
			final = h.middlewares[i](final)
		}
	}

	// 3. Wrap with Group Middlewares (applied before filters to act as inner middlewares)
	if h.group != nil {
		for i := len(h.group.middlewares) - 1; i >= 0; i-- {
			if h.group.middlewares[i] != nil {
				final = h.group.middlewares[i](final)
			}
		}
	}

	// 4. Wrap with Handler-specific Filters (The Gatekeeper Layer)
	// By placing filters here, they become the outermost layer of the compiled handler.
	// This prevents Group and Local middlewares from running if the update is skipped.
	innerChain := final
	final = func(c *Context) {
		ut := c.Update.Type()
		filters := h.typeFilters[ut]
		for _, f := range filters {
			if f != nil && !f(c) {
				c.skipped = true
				return
			}
		}
		innerChain(c)
	}

	h.wrappedHandler = final
}

// runHandler executes the handler and returns true if it wasn't skipped by filters.
// It resets the context's skipped flag before execution.
func runHandler(c *Context, h *Handler) bool {
	c.skipped = false
	h.wrappedHandler(c)
	return !c.skipped
}
