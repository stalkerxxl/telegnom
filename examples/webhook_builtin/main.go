package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/stalkerxxl/telegnom/bot"
	"github.com/stalkerxxl/telegnom/types"
)

// This example demonstrates how to use the built-in webhook server.
// It is suitable for running behind a reverse proxy (Nginx, Traefik) or tunneling tools like ngrok.
func main() {
	// 1. Setup graceful shutdown
	// We want to handle OS signals to gracefully shut down the bot.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 2. Get configuration from environment variables
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	// Webhook URL is the public URL where Telegram will send updates.
	// When using ngrok, this looks like: https://<your-id>.ngrok-free.app/webhook
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("WEBHOOK_URL is required (e.g. https://your-domain.com/webhook)")
	}

	// A secret token used to verify that requests are coming from Telegram.
	// You should generate a random string and keep it secret.
	secretToken := "random-string-for-verification"

	// 3. Initialize the Bot WithWebhookURL automatically parses the URL and
	// configures the internal router to listen on the correct path (e.g.,
	// "/webhook").
	opts := []bot.Option{
		bot.WithWebhookURL(webhookURL),
		bot.WithWebhookSecretToken(secretToken),
		bot.WithPanicRecovery(),
	}

	b, err := bot.NewBot(ctx, token, opts...)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// 4. Register the Webhook in Telegram
	// This makes an API call to Telegram telling them where to send updates.
	log.Printf("Setting webhook to: %s", webhookURL)
	_, err = b.SetWebhook(&bot.SetWebhookParams{
		URL:            webhookURL,
		SecretToken:    secretToken,
		AllowedUpdates: types.AllUpdateTypes(),
	})
	if err != nil {
		log.Fatalf("Failed to set webhook: %v", err)
	}

	// 5. Register Handlers
	b.Handler(func(tg *bot.Context) {
		log.Printf("Received message: %s", tg.Update.Message.Text)
		_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.Update.Message.Chat.ID,
			Text:   "Echo: " + tg.Update.Message.Text,
		})
	}).OnMessage()

	// 6. Start the Built-in Web Server
	// b.StartWebhook starts an HTTP server on the specified address (e.g., ":8080").
	// It internally handles the routing based on WithWebhookURL.
	log.Println("Starting builtin webhook server on :8080...")
	if err := b.StartWebhook(":8080"); err != nil {
		// StartWebhook returns error on context cancellation too, so we check for that
		if !errors.Is(err, context.Canceled) {
			log.Fatalf("Webhook server error: %v", err)
		}
	}

	log.Println("Bot stopped")
}
