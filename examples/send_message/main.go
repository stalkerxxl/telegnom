package main

import (
	"context"
	"log"
	"os"

	"github.com/stalkerxxl/telegnom/bot"
	"github.com/stalkerxxl/telegnom/types"
)

func main() {
	token := os.Getenv("BOT_TOKEN")

	b, _ := bot.NewBot(context.TODO(), token, bot.WithPanicRecovery())
	b.Handler(handleMessage).OnMessage()

	// Удаляем вебхук в Telegram перед выходом
	if _, err := b.DeleteWebhook(nil); err != nil {
		log.Printf("Failed to delete webhook: %v", err)
	}

	b.StartPolling()
}

func handleMessage(c *bot.Context) {
	userID := c.Update.Message.From.ID

	htmlTxt := "<b>Hello World</b>"
	markdownTxt := "*Hello World*"
	markdownV2Txt := "__Hello World__"

	// Send messages with different parse modes
	_, _ = c.Bot.SendMessage(&bot.SendMessageParams{
		ChatID:    userID,
		Text:      htmlTxt + "\nParseMode HTML",
		ParseMode: types.ParseModeHTML,
		ReplyMarkup: &types.ForceReply{
			ForceReply:            true,
			InputFieldPlaceholder: "placeholder",
			Selective:             false,
		},
	})
	//panic("test panic recovery")
	_, _ = c.Bot.SendMessage(&bot.SendMessageParams{
		ChatID:    userID,
		Text:      markdownTxt + "\nParseMode Markdown",
		ParseMode: types.ParseModeMarkdown,
	})
	_, _ = c.Bot.SendMessage(&bot.SendMessageParams{
		ChatID:    userID,
		Text:      markdownV2Txt + "\nParseMode MarkdownV2",
		ParseMode: types.ParseModeMarkdownV2,
	})

	// Send a message with LinkPreviewOptions
	_, _ = c.Bot.SendMessage(&bot.SendMessageParams{
		ChatID: userID,
		Text:   "Check out this link:",
		LinkPreviewOptions: &types.LinkPreviewOptions{
			IsDisabled:       false,
			URL:              "https://www.youtube.com/watch?v=Gg9cNGHl-bg&pp=ygUQbGEgZ3JhbmdlIHp6IHRvcA%3D%3D",
			PreferSmallMedia: false,
			ShowAboveText:    false,
		},
	})

	// Send a message with ReplyParameters
	_, _ = c.Bot.SendMessage(&bot.SendMessageParams{
		ChatID: userID,
		Text:   "This message has reply parameters.",
		ReplyParameters: &types.ReplyParameters{
			MessageID:                c.Update.Message.ID,
			ChatID:                   nil,
			AllowSendingWithoutReply: false,
			Quote:                    c.Update.Message.Text,
			QuoteParseMode:           "",
			QuoteEntities:            nil,
			QuotePosition:            0,
			ChecklistTaskID:          0,
		},
	})

	// ForwardMessage
	_, _ = c.Bot.ForwardMessage(&bot.ForwardMessageParams{
		ChatID:     c.Update.Message.Chat.ID,
		FromChatID: c.Update.Message.Chat.ID,
		MessageID:  c.Update.Message.ID,
	})

	// CopyMessage
	_, _ = c.Bot.CopyMessage(&bot.CopyMessageParams{
		ChatID:     c.Update.Message.Chat.ID,
		FromChatID: c.Update.Message.Chat.ID,
		MessageID:  c.Update.Message.ID,
	})
}
