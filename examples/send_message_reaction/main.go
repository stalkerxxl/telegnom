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

	b.Handler(sendMessageReaction).OnMessage()

	b.StartPolling()
}

func sendMessageReaction(tg *bot.Context) {
	msgID := tg.Update.Message.ID
	reaction1 := &types.ReactionTypeEmoji{Type: "emoji", Emoji: "❤"}
	//reaction2 := &types.ReactionTypeEmoji{Type: "emoji", Emoji: "😍"}

	_, err := tg.Bot.SetMessageReaction(&bot.SetMessageReactionParams{
		MessageID: msgID,
		ChatID:    tg.Update.Message.Chat.ID,
		Reaction:  []types.ReactionType{reaction1},
	})
	if err != nil {
		fmt.Println(err)
	}

}
