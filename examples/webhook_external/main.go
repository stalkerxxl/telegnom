package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"github.com/labstack/echo/v5"
	"github.com/stalkerxxl/telegnom/bot"
	"github.com/stalkerxxl/telegnom/types"
)

// This example demonstrates how to integrate the bot with an external web server (Echo).
// This is useful if you want to run the bot as part of a larger application.
func main() {
	// 1. Setup graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 2. Get configuration
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("WEBHOOK_URL is required")
	}

	secretToken := "external-server-secret"

	// 3. Initialize the Bot
	// Note: We don't strictly need WithWebhookURL here because routing is handled by Echo.
	// However, WithWebhookSecretToken is still needed for verification in WebhookHandler.
	b, err := bot.NewBot(ctx, token,
		bot.WithWebhookSecretToken(secretToken),
		bot.WithPanicRecovery(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Register the Webhook in Telegram
	_, err = b.SetWebhook(&bot.SetWebhookParams{
		URL:            webhookURL,
		SecretToken:    secretToken,
		AllowedUpdates: types.AllUpdateTypes(),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 5. Register Bot Handlers
	b.Handler(func(tg *bot.Context) {
		_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.Update.Message.Chat.ID,
			Text:   "Hello from Echo Server!",
		})
	}).OnMessage()

	// 6. Setup External Web Server (Echo)
	e := echo.New()

	// Register the bot's webhook handler on a specific path
	// b.WebhookHandler() returns a http.HandlerFunc that processes incoming updates
	e.POST("/webhook", echo.WrapHandler(b.WebhookHandler()))

	// 7. Start Bot Workers
	// b.Run() starts the workers that process updates received via the webhook.
	// It is blocking, so we run it in a goroutine.
	go func() {
		log.Println("Starting bot workers...")
		if err := b.Run(); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Printf("Bot error: %v", err)
			}
		}
	}()

	// 8. Start the Echo Server
	go func() {
		log.Println("Starting Echo server on :8080...")
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Wait for interruption signal
	<-ctx.Done()
	log.Println("Shutting down...")

	// 9. Graceful Shutdown
	// We need to stop the web server first to stop receiving new updates.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	// b.Run() will automatically exit when ctx is canceled, so we just wait for the program to exit.
}
