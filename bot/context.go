package bot

import (
	"context"

	"github.com/stalkerxxl/telegnom/types"
)

// Context represents the context of an update being processed by the bot.
type Context struct {
	context.Context
	Bot     *Bot
	Update  *types.Update
	skipped bool // internal flag to indicate if the update was skipped by a handler
}

func newContext(bot *Bot, upd *types.Update) *Context {
	return &Context{
		Context: bot.ctx,
		Bot:     bot,
		Update:  upd,
	}
}

// Set stores a value by creating a new child context.
func (c *Context) Set(key any, value any) {
	c.Context = context.WithValue(c.Context, key, value)
}

// Get retrieves a value from the context chain.
func (c *Context) Get(key any) any {
	return c.Value(key)
}

// EffectiveChat returns the chat where the update happened.
func (c *Context) EffectiveChat() *types.Chat {
	return c.Update.EffectiveChat()
}

// ChatID is a helper that returns the ID of the effective chat, or 0 if not available.
func (c *Context) ChatID() int64 {
	chat := c.EffectiveChat()
	if chat != nil {
		return chat.ID
	}
	return 0
}

// SenderID is a helper that returns the ID of the user who triggered the update, or 0 if not available.
func (c *Context) SenderID() int64 {
	user := c.EffectiveUser()
	if user != nil {
		return user.ID
	}
	return 0
}

// EffectiveUser returns the user who triggered the update.
func (c *Context) EffectiveUser() *types.User {
	return c.Update.EffectiveUser()
}

// EffectiveMessage returns the message from the update.
func (c *Context) EffectiveMessage() *types.Message {
	return c.Update.EffectiveMessage()
}

// Send sends a text message to the effective chat.
func (c *Context) Send(text string) (*types.Message, error) {
	chat := c.EffectiveChat()
	if chat == nil {
		return nil, nil // FIXME возвращать ошибку что EffectiveChat == nil
	}
	return c.Bot.SendMessage(&SendMessageParams{
		ChatID: chat.ID,
		Text:   text,
	})
}

// Reply sends a text message as a reply to the effective message.
func (c *Context) Reply(text string) (*types.Message, error) {
	msg := c.EffectiveMessage()
	if msg == nil {
		return nil, nil // FIXME возвращать ошибку что EffectiveMessage == nil
	}
	return c.Bot.SendMessage(&SendMessageParams{
		ChatID: msg.Chat.ID,
		Text:   text,
		ReplyParameters: &types.ReplyParameters{
			MessageID: msg.ID,
		},
	})
}

// Delete deletes the effective message.
func (c *Context) Delete() error {
	msg := c.EffectiveMessage()
	if msg == nil {
		return nil // FIXME возвращать ошибку что EffectiveMessage == nil
	}
	_, err := c.Bot.DeleteMessage(&DeleteMessageParams{
		ChatID:    msg.Chat.ID,
		MessageID: msg.ID,
	})
	return err
}

// Answer answers the callback query if present in the update.
func (c *Context) Answer(text string) error {
	if c.Update.CallbackQuery == nil {
		return nil // FIXME возвращать ошибку, что нет callback query в обновлении
	}
	_, err := c.Bot.AnswerCallbackQuery(&AnswerCallbackQueryParams{
		CallbackQueryID: c.Update.CallbackQuery.ID,
		Text:            text,
	})
	return err
}
