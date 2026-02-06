package filters

import (
	"regexp"
	"strings"

	"github.com/stalkerxxl/telegnom/bot"
	"github.com/stalkerxxl/telegnom/types"
)

// Simple filters (variables matching bot.Filter signature)

//goland:noinspection GoUnusedExportedFunction
func HasText(c *bot.Context) bool { return c.Update.HasText() }

//goland:noinspection GoUnusedExportedFunction
func HasCaption(c *bot.Context) bool { return c.Update.HasCaption() }

//goland:noinspection GoUnusedExportedFunction
func HasPhoto(c *bot.Context) bool { return c.Update.HasPhoto() }

//goland:noinspection GoUnusedExportedFunction
func HasAudio(c *bot.Context) bool { return c.Update.HasAudio() }

//goland:noinspection GoUnusedExportedFunction
func HasVideo(c *bot.Context) bool { return c.Update.HasVideo() }

//goland:noinspection GoUnusedExportedFunction
func HasDocument(c *bot.Context) bool { return c.Update.HasDocument() }

//goland:noinspection GoUnusedExportedFunction
func HasAnimation(c *bot.Context) bool { return c.Update.HasAnimation() }

//goland:noinspection GoUnusedExportedFunction
func HasVoice(c *bot.Context) bool { return c.Update.HasVoice() }

//goland:noinspection GoUnusedExportedFunction
func HasVideoNote(c *bot.Context) bool { return c.Update.HasVideoNote() }

//goland:noinspection GoUnusedExportedFunction
func HasSticker(c *bot.Context) bool { return c.Update.HasSticker() }

//goland:noinspection GoUnusedExportedFunction
func HasContact(c *bot.Context) bool { return c.Update.HasContact() }

//goland:noinspection GoUnusedExportedFunction
func HasLocation(c *bot.Context) bool { return c.Update.HasLocation() }

//goland:noinspection GoUnusedExportedFunction
func HasVenue(c *bot.Context) bool { return c.Update.HasVenue() }

//goland:noinspection GoUnusedExportedFunction
func HasPoll(c *bot.Context) bool { return c.Update.HasPoll() }

//goland:noinspection GoUnusedExportedFunction
func HasDice(c *bot.Context) bool { return c.Update.HasDice() }

//goland:noinspection GoUnusedExportedFunction
func IsCommand(c *bot.Context) bool { return c.Update.IsCommand() }

//goland:noinspection GoUnusedExportedFunction
func IsReply(c *bot.Context) bool { return c.Update.IsReply() }

//goland:noinspection GoUnusedExportedFunction
func IsForward(c *bot.Context) bool { return c.Update.IsForward() }

//goland:noinspection GoUnusedExportedFunction
func IsTopicMessage(c *bot.Context) bool { return c.Update.IsTopicMessage() }

//goland:noinspection GoUnusedExportedFunction
func IsMediaGroup(c *bot.Context) bool { return c.Update.IsMediaGroup() }

//goland:noinspection GoUnusedExportedFunction
func IsNewChatMember(c *bot.Context) bool { return c.Update.IsNewChatMember() }

//goland:noinspection GoUnusedExportedFunction
func IsLeftChatMember(c *bot.Context) bool { return c.Update.IsLeftChatMember() }

//goland:noinspection GoUnusedExportedFunction
func IsNewChatTitle(c *bot.Context) bool { return c.Update.IsNewChatTitle() }

//goland:noinspection GoUnusedExportedFunction
func IsNewChatPhoto(c *bot.Context) bool { return c.Update.IsNewChatPhoto() }

//goland:noinspection GoUnusedExportedFunction
func IsDeleteChatPhoto(c *bot.Context) bool { return c.Update.IsDeleteChatPhoto() }

//goland:noinspection GoUnusedExportedFunction
func IsPinnedMessage(c *bot.Context) bool { return c.Update.IsPinnedMessage() }

//goland:noinspection GoUnusedExportedFunction
func IsPrivate(c *bot.Context) bool { return c.Update.IsPrivate() }

//goland:noinspection GoUnusedExportedFunction
func IsGroup(c *bot.Context) bool { return c.Update.IsGroup() }

//goland:noinspection GoUnusedExportedFunction
func IsChannel(c *bot.Context) bool { return c.Update.IsChannel() }

// Filter factories (functions returning bot.Filter)

// Command filters messages by a specific command name (e.g. "start").
// It handles commands with or without the leading slash.
//
//goland:noinspection GoUnusedExportedFunction
func Command(name string) bot.Filter {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return func(c *bot.Context) bool {
		msg := c.Update.EffectiveMessage()
		if msg == nil || !msg.IsCommand() {
			return false
		}
		// Split by space to get the command part (e.g. "/start@bot" or just "/start")
		parts := strings.Split(msg.Text, " ")
		cmdPart := parts[0]
		// Handle bot username in command (e.g. "/start@my_bot")
		if strings.Contains(cmdPart, "@") {
			cmdPart = strings.Split(cmdPart, "@")[0]
		}
		return cmdPart == name
	}
}

// Regexp filters message text or caption using a regular expression.
//
//goland:noinspection GoUnusedExportedFunction
func Regexp(pattern string) bot.Filter {
	re := regexp.MustCompile(pattern)
	return func(c *bot.Context) bool {
		msg := c.Update.EffectiveMessage()
		if msg == nil {
			return false
		}
		if msg.Text != "" && re.MatchString(msg.Text) {
			return true
		}
		if msg.Caption != "" && re.MatchString(msg.Caption) {
			return true
		}
		return false
	}
}

// Text filters message text or caption by exact match.
//
//goland:noinspection GoUnusedExportedFunction
func Text(text string) bot.Filter {
	return func(c *bot.Context) bool {
		msg := c.Update.EffectiveMessage()
		if msg == nil {
			return false
		}
		return msg.Text == text || msg.Caption == text
	}
}

// ChatType filters updates by chat type (private, group, supergroup, channel).
//
//goland:noinspection GoUnusedExportedFunction
func ChatType(t types.ChatType) bot.Filter {
	return func(c *bot.Context) bool {
		msg := c.Update.EffectiveMessage()
		if msg == nil {
			// Some updates might not have a message but still have a chat (e.g. callback query)
			if c.Update.CallbackQuery != nil && c.Update.CallbackQuery.Message != nil {
				// We can try to cast CallbackQuery.Message (which is MaybeInaccessibleMessageData)
				if m := c.Update.CallbackQuery.Message.Message(); m != nil {
					return m.Chat.Type == t
				}
			}
			return false
		}
		return msg.Chat.Type == t
	}
}
