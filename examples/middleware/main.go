package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stalkerxxl/telegnom/bot"
	f "github.com/stalkerxxl/telegnom/filters"
)

// This example demonstrates the 3 levels of middleware in Telegnom:
// Global, Group, and Local.
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

	// 1. GLOBAL MIDDLEWARE
	// Runs for EVERY update that reaches the bot.
	b.Use(LoggingMiddleware)

	// 2. GROUP MIDDLEWARE
	// Runs for all handlers within the group.
	admin := b.Group()
	admin.Use(AdminOnlyMiddleware)
	{
		// 3. LOCAL MIDDLEWARE
		// Runs only for this specific handler.
		admin.Handler(func(tg *bot.Context) {
			_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
				ChatID: tg.ChatID(),
				Text:   "Welcome, Master. Here is the secret report.",
			})
		}).OnMessage(f.Command("report")).Use(MetricMiddleware)
	}

	b.Handler(func(tg *bot.Context) {
		_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Hello! Try /report command (needs admin rights).",
		})
	}).OnMessage(f.Command("start"))

	log.Println("Bot with middleware layers started...")
	b.StartPolling()
}

// LoggingMiddleware is a global middleware that logs every update.
func LoggingMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(tg *bot.Context) {
		start := time.Now()
		log.Printf(">>> INCOMING: ID=%d Type=%s", tg.Update.ID, tg.Update.Type())

		next(tg) // Call next middleware/handler

		log.Printf("<<< PROCESSED: ID=%d (took %v)", tg.Update.ID, time.Since(start))
	}
}

// AdminOnlyMiddleware stops the chain if the user is not an admin.
func AdminOnlyMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(tg *bot.Context) {
		// Mock check (in real life use a database or config)
		isAdmin := tg.SenderID() == 123456789

		if !isAdmin {
			_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
				ChatID: tg.ChatID(),
				Text:   "⛔ Access Denied.",
			})
			return // RETURN WITHOUT CALLING next(tg) TO STOP THE CHAIN
		}
		next(tg)
	}
}

// MetricMiddleware is a local middleware for specific handler metrics.
func MetricMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(tg *bot.Context) {
		log.Printf("[METRIC] Executing specialized report handler...")
		next(tg)
	}
}
