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
	b, _ := bot.NewBot(context.TODO(), token)

	b.Handler(sendPhoto).OnMessage(HasPhoto)
	b.Handler(sendVideo).OnMessage(HasVideo)
	b.Handler(sendDocument).OnMessage(HasDocument)

	b.StartPolling()
}

// HasDocument - фильтр наличия документа в сообщении
func HasDocument(c *bot.Context) bool {
	return c.Update.Message.Document != nil
}

// HasPhoto - фильтр наличия фото в сообщении
func HasPhoto(c *bot.Context) bool {
	return c.Update.Message.Photo != nil
}

// HasVideo - фильтр наличия видео в сообщении
func HasVideo(c *bot.Context) bool {
	return c.Update.Message.Video != nil
}

func sendPhoto(c *bot.Context) {
	var photoPath = "./.files/photo.jpg"
	//var photoURL = "https://bun.uptrace.dev/bun/cover.png"
	file, _ := os.Open(photoPath)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	fmt.Println("sendPhoto handler")

	if photo := c.Update.Message.Photo.Last(); photo != nil {
		// Отправка photo по file_id
		_, err := c.Bot.SendPhoto(&bot.SendPhotoParams{
			ChatID: c.Update.Message.Chat.ID,
			Photo:  &types.InputFile{ID: photo.FileID},
			//Photo:   &types.InputFile{Path: photoPath},
			//Photo: &types.InputFile{URL: photoURL},
			//Photo: &types.InputFile{Reader: file, Name: "photo_name.jpg"},
		})
		if err != nil {
			fmt.Printf("SendPhoto err: %v\n", err)
		}
	}
}

func sendVideo(c *bot.Context) {
	msg, _ := c.Bot.SendVideo(&bot.SendVideoParams{
		ChatID: 198172266,
		Video: &types.InputFile{
			ID: c.Update.Message.Video.FileID,
			//Path: "./.files/video.mp4",
			//URL:      video_url,
			//Reader:   video_reader,
			//Name: "",
		},
		Cover: &types.InputFile{
			Path: "./.files/cover.jpg",
		},
		//Thumbnail: &types.InputFile{
		//	Path: "./.files/thumb.png",
		//},
	})
	fmt.Printf("%#v", msg.Video)
}

func sendDocument(c *bot.Context) {
	inputDoc := c.Update.Message.Document
	_, err := c.Bot.SendDocument(&bot.SendDocumentParams{
		ChatID: c.Update.Message.Chat.ID,
		Document: &types.InputFile{
			ID: inputDoc.FileID,
		},
	})
	if err != nil {
		fmt.Printf("SendDocument err: %v\n", err)
	}

	reader := strings.NewReader("Контент файла прямо из памяти")
	_, err = c.Bot.SendDocument(&bot.SendDocumentParams{
		ChatID: c.Update.Message.Chat.ID,
		Document: &types.InputFile{
			Reader: reader,
			Name:   "my_file.txt",
		},
	})
	if err != nil {
		fmt.Printf("SendDocument err: %v\n", err)
	}

}

func sendAudio(c *bot.Context) {}

func sendAnimation(c *bot.Context) {}
