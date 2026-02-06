// Package types contains type definitions for working with the Telegram Bot API.
// It includes structures for updates, messages, and other entities.
package types

// Update represents an incoming update.
//
//	‼️ At most one of the optional parameters can be present in any given update.
//
// See https://core.telegram.org/bots/api#update
type Update struct {
	ID                      int64                        `json:"update_id"`
	Message                 *Message                     `json:"message,omitempty"`
	EditedMessage           *Message                     `json:"edited_message,omitempty"`
	ChannelPost             *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	Poll                    *Poll                        `json:"poll,omitempty"`
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
}

type UpdateType string

type AllowedUpdates []UpdateType

// allowed_updates https://core.telegram.org/bots/api#update
const (
	MessageType                 UpdateType = "message"
	EditedMessageType           UpdateType = "edited_message"
	ChannelPostType             UpdateType = "channel_post"
	EditedChannelPostType       UpdateType = "edited_channel_post"
	BusinessConnectionType      UpdateType = "business_connection"
	BusinessMessageType         UpdateType = "business_message"
	EditedBusinessMessageType   UpdateType = "edited_business_message"
	DeletedBusinessMessagesType UpdateType = "deleted_business_messages"
	MessageReactionType         UpdateType = "message_reaction"
	MessageReactionCountType    UpdateType = "message_reaction_count"
	InlineQueryType             UpdateType = "inline_query"
	ChosenInlineResultType      UpdateType = "chosen_inline_result"
	CallbackQueryType           UpdateType = "callback_query"
	ShippingQueryType           UpdateType = "shipping_query"
	PreCheckoutQueryType        UpdateType = "pre_checkout_query"
	PurchasedPaidMediaType      UpdateType = "purchased_paid_media"
	PollType                    UpdateType = "poll"
	PollAnswerType              UpdateType = "poll_answer"
	MyChatMemberType            UpdateType = "my_chat_member"
	ChatMemberType              UpdateType = "chat_member"
	ChatJoinRequestType         UpdateType = "chat_join_request"
	ChatBoostType               UpdateType = "chat_boost"
	RemovedChatBoostType        UpdateType = "removed_chat_boost"
)

// Type returns the UpdateType of the update by identifying the non-nil field
// among its components. Returns an empty string if none.
//
//goland:noinspection DuplicatedCode
func (u *Update) Type() UpdateType {
	if u.Message != nil {
		return MessageType
	}
	if u.EditedMessage != nil {
		return EditedMessageType
	}
	if u.ChannelPost != nil {
		return ChannelPostType
	}
	if u.EditedChannelPost != nil {
		return EditedChannelPostType
	}
	if u.BusinessConnection != nil {
		return BusinessConnectionType
	}
	if u.BusinessMessage != nil {
		return BusinessMessageType
	}
	if u.EditedBusinessMessage != nil {
		return EditedBusinessMessageType
	}
	if u.DeletedBusinessMessages != nil {
		return DeletedBusinessMessagesType
	}
	if u.MessageReaction != nil {
		return MessageReactionType
	}
	if u.MessageReactionCount != nil {
		return MessageReactionCountType
	}
	if u.InlineQuery != nil {
		return InlineQueryType
	}
	if u.ChosenInlineResult != nil {
		return ChosenInlineResultType
	}
	if u.CallbackQuery != nil {
		return CallbackQueryType
	}
	if u.ShippingQuery != nil {
		return ShippingQueryType
	}
	if u.PreCheckoutQuery != nil {
		return PreCheckoutQueryType
	}
	if u.PurchasedPaidMedia != nil {
		return PurchasedPaidMediaType
	}
	if u.Poll != nil {
		return PollType
	}
	if u.PollAnswer != nil {
		return PollAnswerType
	}
	if u.MyChatMember != nil {
		return MyChatMemberType
	}
	if u.ChatMember != nil {
		return ChatMemberType
	}
	if u.ChatJoinRequest != nil {
		return ChatJoinRequestType
	}
	if u.ChatBoost != nil {
		return ChatBoostType
	}
	if u.RemovedChatBoost != nil {
		return RemovedChatBoostType
	}
	return ""
}

// AllUpdateTypes returns a slice of all possible UpdateType values. This can be
// useful for bot.WithAllowedUpdates option.
func AllUpdateTypes() []UpdateType {
	return []UpdateType{
		MessageType,
		EditedMessageType,
		ChannelPostType,
		EditedChannelPostType,
		BusinessConnectionType,
		BusinessMessageType,
		EditedBusinessMessageType,
		DeletedBusinessMessagesType,
		MessageReactionType,
		MessageReactionCountType,
		InlineQueryType,
		ChosenInlineResultType,
		CallbackQueryType,
		ShippingQueryType,
		PreCheckoutQueryType,
		PurchasedPaidMediaType,
		PollType,
		PollAnswerType,
		MyChatMemberType,
		ChatMemberType,
		ChatJoinRequestType,
		ChatBoostType,
		RemovedChatBoostType,
	}
}

// Proxy methods that safely check the property in the EffectiveMessage

// EffectiveMessage returns the available message object in the update.
// It checks Message, EditedMessage, ChannelPost, EditedChannelPost, BusinessMessage,
// and EditedBusinessMessage in that order.
func (u *Update) EffectiveMessage() *Message {
	if u.Message != nil {
		return u.Message
	}
	if u.EditedMessage != nil {
		return u.EditedMessage
	}
	if u.ChannelPost != nil {
		return u.ChannelPost
	}
	if u.EditedChannelPost != nil {
		return u.EditedChannelPost
	}
	if u.BusinessMessage != nil {
		return u.BusinessMessage
	}
	if u.EditedBusinessMessage != nil {
		return u.EditedBusinessMessage
	}
	return nil
}

func (u *Update) HasText() bool      { m := u.EffectiveMessage(); return m != nil && m.HasText() }
func (u *Update) HasCaption() bool   { m := u.EffectiveMessage(); return m != nil && m.HasCaption() }
func (u *Update) HasPhoto() bool     { m := u.EffectiveMessage(); return m != nil && m.HasPhoto() }
func (u *Update) HasAudio() bool     { m := u.EffectiveMessage(); return m != nil && m.HasAudio() }
func (u *Update) HasVideo() bool     { m := u.EffectiveMessage(); return m != nil && m.HasVideo() }
func (u *Update) HasDocument() bool  { m := u.EffectiveMessage(); return m != nil && m.HasDocument() }
func (u *Update) HasAnimation() bool { m := u.EffectiveMessage(); return m != nil && m.HasAnimation() }
func (u *Update) HasVoice() bool     { m := u.EffectiveMessage(); return m != nil && m.HasVoice() }
func (u *Update) HasVideoNote() bool { m := u.EffectiveMessage(); return m != nil && m.HasVideoNote() }
func (u *Update) HasSticker() bool   { m := u.EffectiveMessage(); return m != nil && m.HasSticker() }
func (u *Update) HasContact() bool   { m := u.EffectiveMessage(); return m != nil && m.HasContact() }
func (u *Update) HasLocation() bool  { m := u.EffectiveMessage(); return m != nil && m.HasLocation() }
func (u *Update) HasVenue() bool     { m := u.EffectiveMessage(); return m != nil && m.HasVenue() }
func (u *Update) HasPoll() bool      { m := u.EffectiveMessage(); return m != nil && m.HasPoll() }

func (u *Update) HasDice() bool        { m := u.EffectiveMessage(); return m != nil && m.HasDice() }
func (u *Update) IsCommand() bool      { m := u.EffectiveMessage(); return m != nil && m.IsCommand() }
func (u *Update) IsReply() bool        { m := u.EffectiveMessage(); return m != nil && m.IsReply() }
func (u *Update) IsForward() bool      { m := u.EffectiveMessage(); return m != nil && m.IsForward() }
func (u *Update) IsTopicMessage() bool { m := u.EffectiveMessage(); return m != nil && m.InTopic() }

func (u *Update) IsMediaGroup() bool { m := u.EffectiveMessage(); return m != nil && m.IsMediaGroup() }
func (u *Update) IsNewChatMember() bool {
	m := u.EffectiveMessage()
	return m != nil && m.IsNewChatMember()
}
func (u *Update) IsLeftChatMember() bool {
	m := u.EffectiveMessage()
	return m != nil && m.IsLeftChatMember()
}
func (u *Update) IsNewChatTitle() bool {
	m := u.EffectiveMessage()
	return m != nil && m.IsNewChatTitle()
}
func (u *Update) IsNewChatPhoto() bool {
	m := u.EffectiveMessage()
	return m != nil && m.IsNewChatPhoto()
}
func (u *Update) IsDeleteChatPhoto() bool {
	m := u.EffectiveMessage()
	return m != nil && m.IsDeleteChatPhoto()
}

func (u *Update) IsPinnedMessage() bool {
	m := u.EffectiveMessage()
	return m != nil && m.IsPinnedMessage()
}
func (u *Update) IsPrivate() bool { m := u.EffectiveMessage(); return m != nil && m.IsPrivate() }
func (u *Update) IsGroup() bool   { m := u.EffectiveMessage(); return m != nil && m.IsGroup() }

// Special Update level checks

func (u *Update) IsChannel() bool     { m := u.EffectiveMessage(); return m != nil && m.IsChannel() }
func (u *Update) IsCallback() bool    { return u.CallbackQuery != nil }
func (u *Update) IsInlineQuery() bool { return u.InlineQuery != nil }

func (u *Update) IsEdited() bool {
	return u.EditedMessage != nil || u.EditedChannelPost != nil || u.EditedBusinessMessage != nil
}

// EffectiveChat returns the chat associated with the update.
func (u *Update) EffectiveChat() *Chat {
	if m := u.EffectiveMessage(); m != nil {
		return m.Chat
	}
	if u.CallbackQuery != nil {
		return u.CallbackQuery.EffectiveChat()
	}
	if u.MyChatMember != nil {
		return u.MyChatMember.EffectiveChat()
	}
	if u.ChatMember != nil {
		return u.ChatMember.EffectiveChat()
	}
	if u.ChatJoinRequest != nil {
		return u.ChatJoinRequest.EffectiveChat()
	}
	return nil
}

// EffectiveUser returns the user associated with the update.
func (u *Update) EffectiveUser() *User {
	if m := u.EffectiveMessage(); m != nil {
		return m.From
	}
	if u.CallbackQuery != nil {
		return u.CallbackQuery.EffectiveUser()
	}
	if u.InlineQuery != nil {
		return u.InlineQuery.EffectiveUser()
	}
	if u.ChosenInlineResult != nil {
		return u.ChosenInlineResult.EffectiveUser()
	}
	if u.ShippingQuery != nil {
		return u.ShippingQuery.EffectiveUser()
	}
	if u.PreCheckoutQuery != nil {
		return u.PreCheckoutQuery.EffectiveUser()
	}
	if u.MyChatMember != nil {
		return u.MyChatMember.EffectiveUser()
	}
	if u.ChatMember != nil {
		return u.ChatMember.EffectiveUser()
	}
	if u.ChatJoinRequest != nil {
		return u.ChatJoinRequest.EffectiveUser()
	}
	return nil
}
