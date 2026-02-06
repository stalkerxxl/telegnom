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

// This example demonstrates how to use Handler Groups to organize handlers
// and apply shared middleware (like authentication or logging) to them.
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

	// 1. Public Group (Default)
	// You don't always need a group, but you can create one for logical separation.
	public := b.Group()
	{
		public.Handler(func(tg *bot.Context) {
			_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
				ChatID: tg.ChatID(),
				Text:   "Welcome to the public part of the bot!",
			})
		}).OnMessage(f.Command("start"))
	}

	// 2. Protected/Admin Group
	// This group has a middleware that checks if the user is an admin.
	// This middleware will run for ALL handlers in this group.
	admin := b.Group()
	admin.Use(AdminOnlyMiddleware) // Apply shared behavior
	{
		admin.Handler(func(tg *bot.Context) {
			_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
				ChatID: tg.ChatID(),
				Text:   "Admin stats: 1000 users active.",
			})
		}).OnMessage(f.Command("stats"))

		admin.Handler(func(tg *bot.Context) {
			_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
				ChatID: tg.ChatID(),
				Text:   "All systems green.",
			})
		}).OnMessage(f.Command("status"))
	}

	log.Println("Bot with groups started...")
	b.StartPolling()
}

// AdminOnlyMiddleware checks if the user ID matches the admin ID.
// If not, it stops the chain and sends an error message.
func AdminOnlyMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	const AdminID = 123456789 // Replace with real admin ID

	return func(tg *bot.Context) {
		if tg.SenderID() != AdminID {
			log.Printf("Unauthorized access attempt by %d", tg.SenderID())
			_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
				ChatID: tg.ChatID(),
				Text:   "⛔ Access denied. Admin only.",
			})
			return // Stop execution here! The handler won't be called.
		}

		// If user is admin, continue to the next middleware or handler.
		next(tg)
	}
}
