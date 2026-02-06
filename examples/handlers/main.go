package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/stalkerxxl/telegnom/bot"
	f "github.com/stalkerxxl/telegnom/filters"
)

// This example demonstrates advanced handler configurations in Telegnom.
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

	// 1. Basic Handler
	// Simple registration for a specific update type (Message)
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Hello! I am a basic handler.",
		})
	}).OnMessage(f.Command("start"))

	// 2. Multi-Trigger Handler
	// One function handling multiple types of updates (Message and EditedMessage).
	// Useful for logic that should work for both new and edited messages.
	b.Handler(func(tg *bot.Context) {
		msg := tg.Update.EffectiveMessage()
		text := "I see a text message"
		if tg.Update.IsEdited() {
			text = "I see an EDITED text message"
		}

		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: msg.Chat.ID,
			Text:   text + ": " + msg.Text,
		})
	}).OnMessage(f.HasText).OnEditedMessage(f.HasText)

	// 3. Handler with Local Middleware
	// You can add middleware that applies ONLY to this specific handler.
	// This is great for one-off logging, rate limiting, or specific checks.
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Secret data revealed!",
		})
	}).OnMessage(f.Command("secret")).Use(secretLoggerMiddleware)

	// 4. Handler Closure (Dependency Injection)
	// Passing external dependencies (like a database or config) to a handler
	// without using global variables.
	db := &MockDatabase{}
	b.Handler(userInfoHandler(db)).OnMessage(f.Command("info"))

	// 5. Handling different update types
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Button clicked!",
		})
		// Always answer callback queries to stop the loading animation!
		tg.Bot.AnswerCallbackQuery(&bot.AnswerCallbackQueryParams{
			CallbackQueryID: tg.Update.CallbackQuery.ID,
			Text:            "Done",
		})
	}).OnCallbackQuery()

	log.Println("Bot started...")
	b.StartPolling()
}

// secretLoggerMiddleware is a local middleware example
func secretLoggerMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(tg *bot.Context) {
		log.Printf("[SECRET ACCESS] User %d tried to access secret", tg.SenderID())
		next(tg)
	}
}

// MockDatabase simulates a database connection
type MockDatabase struct{}

func (db *MockDatabase) GetUserStatus(userID int64) string {
	return "Premium"
}

// userInfoHandler returns a handler function that closes over the 'db' variable.
func userInfoHandler(db *MockDatabase) bot.HandlerFunc {
	return func(tg *bot.Context) {
		status := db.GetUserStatus(tg.SenderID())
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Your status is: " + status,
		})
	}
}
