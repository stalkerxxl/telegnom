package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stalkerxxl/telegnom/bot"
	"github.com/stalkerxxl/telegnom/types"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	b, _ := bot.NewBot(context.TODO(), token)

	b.Handler(sendPoll).OnMessage()

	_, _ = b.DeleteWebhook(nil)

	b.StartPolling()
}

func sendPoll(tg *bot.Context) {
	options := []types.InputPollOption{
		{Text: "Option 1"},
		{Text: "Option 2"},
	}

	_, err := tg.Bot.SendPoll(
		&bot.SendPollParams{
			ChatID:   tg.Update.Message.Chat.ID,
			Question: "Question text",
			Options:  options,
		})
	if err != nil {
		fmt.Println(err)
	}
}
