package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/stalkerxxl/telegnom/bot"
	f "github.com/stalkerxxl/telegnom/filters"
)

// ---------------------------------------------------------
// 1. Service Layer (Business Logic)
// ---------------------------------------------------------

type UserService struct {
	// Here you would have a database connection or other clients
}

func (s *UserService) GetProfile(userID int64) string {
	return fmt.Sprintf("User profile for ID %d: [Active]", userID)
}

func (s *UserService) UpdateBalance(userID int64, amount int) error {
	log.Printf("Updating balance for user %d by %d", userID, amount)
	return nil
}

// ---------------------------------------------------------
// 2. Handler Layer (Controller)
// ---------------------------------------------------------

// UserHandler groups all user-related bot handlers.
type UserHandler struct {
	service *UserService
}

// NewUserHandler creates a new handler instance with injected dependencies.
func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Profile is a handler method for the /profile command.
func (h *UserHandler) Profile(tg *bot.Context) {
	profile := h.service.GetProfile(tg.SenderID())
	_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
		ChatID: tg.ChatID(),
		Text:   profile,
	})
}

// TopUp is a handler method for the /topup command.
func (h *UserHandler) TopUp(tg *bot.Context) {
	_ = h.service.UpdateBalance(tg.SenderID(), 100)
	_, _ = tg.Bot.SendMessage(&bot.SendMessageParams{
		ChatID: tg.ChatID(),
		Text:   "Balance topped up by 100!",
	})
}

// ---------------------------------------------------------
// 3. Main Setup
// ---------------------------------------------------------

func runStructHandlerExample() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	b, err := bot.NewBot(ctx, token)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize services and handlers
	userService := &UserService{}
	userHandler := NewUserHandler(userService)

	// Register methods as handlers.
	// In Go, instance methods can be passed as values: userHandler.Profile
	userGroup := b.Group()
	{
		userGroup.Handler(userHandler.Profile).OnMessage(f.Command("profile"))
		userGroup.Handler(userHandler.TopUp).OnMessage(f.Command("topup"))
	}

	log.Println("Bot with struct-based handlers started...")
	b.StartPolling()
}
