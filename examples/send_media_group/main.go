package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/stalkerxxl/telegnom/bot"
	"github.com/stalkerxxl/telegnom/types"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	b, _ := bot.NewBot(context.TODO(), "123:token")

	//b.Handler(sendMediaGroup).OnMessage() // FIXME добавить фильтры
	b.Handler(sendDocumentGroup).OnMessage()

	b.StartPolling()

}

func sendMediaGroup(c *bot.Context) {
	videoPath := "./.files/video.mp4"
	photoPath := "./.files/photo.jpg"
	coverPath := "./.files/cover.jpg"
	thumbPath := "./.files/thumb.png"

	video := &types.InputMediaVideo{
		Type:           "video", // FIXME сделать автозаполнение поля?
		Media:          &types.InputFile{Path: videoPath},
		Thumbnail:      &types.InputFile{Path: thumbPath},
		Cover:          &types.InputFile{Path: coverPath},
		StartTimestamp: 0,
		Caption:        "Caption for video",
	}

	photo := &types.InputMediaPhoto{
		Type:    "photo",
		Media:   &types.InputFile{Path: photoPath},
		Caption: "Caption for photo",
	}

	var mediaGroup []types.InputMedia
	mediaGroup = append(mediaGroup, video, photo)

	_, err := c.Bot.SendMediaGroup(&bot.SendMediaGroupParams{
		ChatID: c.Update.Message.Chat.ID,
		Media:  mediaGroup,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func sendDocumentGroup(c *bot.Context) {
	reader1 := strings.NewReader("Контент файла прямо из памяти")
	reader2 := strings.NewReader("Контент 2-го файла прямо из памяти")
	doc1 := &types.InputMediaDocument{
		Type: "document",
		Media: &types.InputFile{
			Reader: reader1,
			Name:   "my_file.txt",
		},
	}

	doc2 := &types.InputMediaDocument{
		Type: "document",
		Media: &types.InputFile{
			Reader: reader2,
			Name:   "my_file2.txt",
		},
	}

	docGroup := []types.InputMedia{doc1, doc2}
	_, err := c.Bot.SendMediaGroup(&bot.SendMediaGroupParams{
		ChatID: c.Update.Message.Chat.ID,
		Media:  docGroup,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
