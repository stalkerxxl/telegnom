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

	b, _ := bot.NewBot(context.TODO(), "123:token")
	b.Handler(handleMessage).OnMessage()

	b.StartPolling()
}

func handleMessage(tg *bot.Context) {
	videoPath := "./.files/video.mp4"
	//photoPath := "./.files/photo.jpg"
	coverPath := "./.files/cover.jpg"
	thumbPath := "./.files/thumb.png"

	video := &types.InputPaidMediaVideo{
		Type:      "video", // FIXME сделать автозаполнение поля?
		Media:     &types.InputFile{Path: videoPath},
		Thumbnail: &types.InputFile{Path: thumbPath},
		Cover:     &types.InputFile{Path: coverPath},
	}

	var paidGroup []types.InputPaidMedia
	paidGroup = append(paidGroup, video)

	_, err := tg.Bot.SendPaidMedia(&bot.SendPaidMediaParams{
		ChatID:    tg.Update.Message.Chat.ID,
		Media:     paidGroup,
		StarCount: 1,
	})
	if err != nil {
		fmt.Println(err)
	}
}
