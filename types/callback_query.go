package types

// CallbackQuery represents an incoming callback query from a callback button in
// an inline keyboard. If the button that originated the query was attached to a
// message sent by the bot, the field message will be present. If the button was
// attached to a message sent via the bot (in inline mode), the field
// inline_message_id will be present. Exactly one of the fields data or
// game_short_name will be present.
//
//	NOTE: After the user presses a callback button, Telegram clients will display
//	a progress bar until you call AnswerCallbackQuery. It is, therefore, necessary
//	to react by calling AnswerCallbackQuery even if no notification to the user is
//	needed (e.g., without specifying any of the optional parameters).
//
// See https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string                        `json:"id"`
	From            *User                         `json:"from"`
	Message         *MaybeInaccessibleMessageData `json:"message,omitempty"`
	InlineMessageID string                        `json:"inline_message_id,omitempty"`
	ChatInstance    string                        `json:"chat_instance,omitempty"`
	Data            string                        `json:"data,omitempty"`
	GameShortName   string                        `json:"game_short_name,omitempty"`
}

// EffectiveChat returns the chat of the message if present.
func (c *CallbackQuery) EffectiveChat() *Chat {
	if c.Message != nil {
		if m := c.Message.Message(); m != nil {
			return m.Chat
		}
		if im := c.Message.InaccessibleMessage(); im != nil {
			return im.Chat
		}
	}
	return nil
}

// EffectiveUser returns the user who sent the callback query.
func (c *CallbackQuery) EffectiveUser() *User {
	return c.From
}
