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

	b.Handler(setChatPhoto).OnEditedMessage()

	b.StartPolling()
}

func setChatPhoto(tg *bot.Context) {
	chatID := -1003012322056
	var photoPath = "./.files/photo.jpg"
	file, _ := os.Open(photoPath)
	defer file.Close()

	_, err := tg.Bot.SetChatPhoto(&bot.SetChatPhotoParams{
		ChatID: chatID,
		Photo:  &types.InputFile{Reader: file},
	})
	if err != nil {
		fmt.Println(err)
	}
}
