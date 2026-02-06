package types

// HasText returns true if the message has text.
func (m *Message) HasText() bool {
	return m.Text != ""
}

// HasCaption returns true if the message has a caption.
func (m *Message) HasCaption() bool {
	return m.Caption != ""
}

// HasPhoto returns true if the message has photos.
func (m *Message) HasPhoto() bool {
	return len(m.Photo) > 0
}

// HasAudio returns true if the message has audio.
func (m *Message) HasAudio() bool {
	return m.Audio != nil
}

// HasVideo returns true if the message has video.
func (m *Message) HasVideo() bool {
	return m.Video != nil
}

// HasDocument returns true if the message has a document.
func (m *Message) HasDocument() bool {
	return m.Document != nil
}

// HasAnimation returns true if the message has an animation.
func (m *Message) HasAnimation() bool {
	return m.Animation != nil
}

// HasVoice returns true if the message has a voice note.
func (m *Message) HasVoice() bool {
	return m.Voice != nil
}

// HasVideoNote returns true if the message has a video note.
func (m *Message) HasVideoNote() bool {
	return m.VideoNote != nil
}

// HasSticker returns true if the message has a sticker.
func (m *Message) HasSticker() bool {
	return m.Sticker != nil
}

// HasContact returns true if the message has a contact.
func (m *Message) HasContact() bool {
	return m.Contact != nil
}

// HasLocation returns true if the message has a location.
func (m *Message) HasLocation() bool {
	return m.Location != nil
}

// HasVenue returns true if the message has a venue.
func (m *Message) HasVenue() bool {
	return m.Venue != nil
}

// HasPoll returns true if the message has a poll.
func (m *Message) HasPoll() bool {
	return m.Poll != nil
}

// HasDice returns true if the message has dice/casino animation.
func (m *Message) HasDice() bool {
	return m.Dice != nil
}

// IsCommand returns true if the message is a bot command.
func (m *Message) IsCommand() bool {
	if m.Text == "" {
		return false
	}
	for _, entity := range m.Entities {
		if entity.Type == "bot_command" && entity.Offset == 0 {
			return true
		}
	}
	return false
}

// IsReply returns true if the message is a reply to another message.
func (m *Message) IsReply() bool {
	return m.ReplyToMessage != nil
}

// IsForward returns true if the message is a forwarded message.
func (m *Message) IsForward() bool {
	return m.ForwardOrigin != nil
}

// InTopic returns true if the message is sent to a forum topic.
func (m *Message) InTopic() bool {
	return m.IsTopicMessage
}

// IsMediaGroup returns true if the message is part of a media group (album).
func (m *Message) IsMediaGroup() bool {
	return m.MediaGroupID != ""
}

// IsPrivate returns true if the message is from a private chat.
func (m *Message) IsPrivate() bool {
	return m.Chat.Type == "private"
}

// IsGroup returns true if the message is from a group or supergroup.
func (m *Message) IsGroup() bool {
	return m.Chat.Type == "group" || m.Chat.Type == "supergroup"
}

// IsChannel returns true if the message is from a channel.
func (m *Message) IsChannel() bool {
	return m.Chat.Type == "channel"
}

// IsNewChatMember returns true if the message indicates new members joined the chat.
func (m *Message) IsNewChatMember() bool {
	return len(m.NewChatMembers) > 0
}

// IsLeftChatMember returns true if the message indicates a member left the chat.
func (m *Message) IsLeftChatMember() bool {
	return m.LeftChatMember != nil
}

// IsNewChatTitle returns true if the chat title was changed.
func (m *Message) IsNewChatTitle() bool {
	return m.NewChatTitle != ""
}

// IsNewChatPhoto returns true if the chat photo was changed.
func (m *Message) IsNewChatPhoto() bool {
	return len(m.NewChatPhoto) > 0
}

// IsDeleteChatPhoto returns true if the chat photo was deleted.
func (m *Message) IsDeleteChatPhoto() bool {
	return m.DeleteChatPhoto
}

// IsPinnedMessage returns true if a message was pinned.
func (m *Message) IsPinnedMessage() bool {
	return m.PinnedMessage != nil
}
