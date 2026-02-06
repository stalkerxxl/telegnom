package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/stalkerxxl/telegnom/bot"
	f "github.com/stalkerxxl/telegnom/filters"
	_ "github.com/stalkerxxl/telegnom/types"
)

// This example demonstrates the powerful filtering system of Telegnom.
// Filters determine IF a handler should be executed.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	b, err := bot.NewBot(ctx, token, bot.WithPanicRecovery())
	if err != nil {
		log.Fatal(err)
	}

	// 1. Package Filters (High-level)
	// Most common use case. 'filters' package contains ready-to-use filters.
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "You sent the /start command!",
		})
	}).OnMessage(f.Command("start"))

	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Nice photo!",
		})
	}).OnMessage(f.HasPhoto)

	// 2. Custom Filter (Function)
	// You can write your own logic.
	// Let's create a filter that only accepts messages with the word "telegnom" (case-insensitive)
	isTelegnomMentioned := func(tg *bot.Context) bool {
		if !tg.Update.HasText() {
			return false
		}
		text := strings.ToLower(tg.Update.Message.Text)
		return strings.Contains(text, "telegnom")
	}

	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "I heard my name!",
		})
	}).OnMessage(isTelegnomMentioned)

	// 3. Combining Filters (AND logic)
	// When you pass multiple filters to OnMessage, they all must return true.
	// This handler only runs if it's a Private chat AND has a Document.
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Received a document in private chat.",
		})
	}).OnMessage(f.IsPrivate, f.HasDocument)

	// 4. Using types.Update methods inside a handler (Manual filtering)
	// Sometimes you want to filter INSIDE the handler to have branching logic.
	b.Handler(func(tg *bot.Context) {
		// Using the types.Update helper method
		if tg.Update.IsForward() {
			tg.Bot.SendMessage(&bot.SendMessageParams{
				ChatID: tg.ChatID(),
				Text:   "This is a forwarded message.",
			})
			return
		}
		// Using the types.Message helper method
		msg := tg.Update.EffectiveMessage()
		if msg.IsReply() {
			tg.Bot.SendMessage(&bot.SendMessageParams{
				ChatID: tg.ChatID(),
				Text:   "This is a reply to another message.",
			})
			return
		}

		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Just a normal message.",
		})
	}).OnMessage(f.HasText) // We filter for Text first to ensure EffectiveMessage is not nil

	log.Println("Bot started...")
	b.StartPolling()
}
