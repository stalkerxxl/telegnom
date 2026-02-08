package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stalkerxxl/telegnom/bot"
)

func main() {
	token := os.Getenv("BOT_TOKEN")

	b, _ := bot.NewBot(context.TODO(), token)

	// Пример вызова метода GetMe
	me, err := b.GetMe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", me)

	// MyDescription
	ok, err := b.SetMyShortDescription(
		&bot.SetMyShortDescriptionParams{
			ShortDescription: "A short description",
			LanguageCode:     "ru",
		},
	)
	fmt.Printf("OK: %#v, ERROR: %#v\n", ok, err)
	desc, err := b.GetMyDescription(&bot.GetMyDescriptionParams{LanguageCode: "ru"})
	fmt.Printf("DESC: %#v, ERROR: %#v\n", desc, err)

	// BotCommandScope
	// TODO
}
