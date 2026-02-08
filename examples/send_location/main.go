package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stalkerxxl/telegnom/bot"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	b, _ := bot.NewBot(context.TODO(), "123:token")

	b.Handler(sendLocation).OnMessage()

	b.StartPolling()
}

func sendLocation(tg *bot.Context) {
	if tg.Update.Message.Location != nil {
		_, err := tg.Bot.SendLocation(&bot.SendLocationParams{
			ChatID:    tg.Update.Message.Chat.ID,
			Latitude:  tg.Update.Message.Location.Latitude,
			Longitude: tg.Update.Message.Location.Longitude,
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}
