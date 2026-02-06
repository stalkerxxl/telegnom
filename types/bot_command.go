package types

import (
	"encoding/json"
	"fmt"
)

// BotCommand represents a bot command.
//
// See https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// BotCommandScope represents the scope to which bot commands are applied.
// Currently, the following 7 scopes are supported: BotCommandScopeDefault,
// BotCommandScopeAllPrivateChats, BotCommandScopeAllGroupChats,
// BotCommandScopeAllChatAdministrators, BotCommandScopeChat,
// BotCommandScopeChatAdministrators, BotCommandScopeChatMember.
//
// See https://core.telegram.org/bots/api#botcommandscope
type BotCommandScope interface {
	isBotCommandScope()
}

// BotCommandScopeData is a wrapper for the BotCommandScope interface.
//
// See https://core.telegram.org/bots/api#botcommandscope
type BotCommandScopeData struct {
	impl BotCommandScope
}

func (bcs *BotCommandScopeData) MarshalJSON() ([]byte, error) {
	if bcs.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(bcs.impl)
}

func (bcs *BotCommandScopeData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "default":
		var val BotCommandScopeDefault
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bcs.impl = &val
	case "all_private_chats":
		var val BotCommandScopeAllPrivateChats
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bcs.impl = &val
	case "all_group_chats":
		var val BotCommandScopeAllGroupChats
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bcs.impl = &val
	case "all_chat_administrators":
		var val BotCommandScopeAllChatAdministrators
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bcs.impl = &val
	case "chat":
		var val BotCommandScopeChat
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bcs.impl = &val
	case "chat_administrators":
		var val BotCommandScopeChatAdministrators
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bcs.impl = &val
	case "chat_member":
		var val BotCommandScopeChatMember
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bcs.impl = &val
	default:
		return fmt.Errorf("unknown bot_command_scope_type: %s", helper.Type)
	}

	return nil
}

func (bcs *BotCommandScopeData) Default() *BotCommandScopeDefault {
	if val, ok := bcs.impl.(*BotCommandScopeDefault); ok {
		return val
	}
	return nil
}

func (bcs *BotCommandScopeData) AllPrivateChats() *BotCommandScopeAllPrivateChats {
	if val, ok := bcs.impl.(*BotCommandScopeAllPrivateChats); ok {
		return val
	}
	return nil
}

func (bcs *BotCommandScopeData) AllGroupChats() *BotCommandScopeAllGroupChats {
	if val, ok := bcs.impl.(*BotCommandScopeAllGroupChats); ok {
		return val
	}
	return nil
}

func (bcs *BotCommandScopeData) AllChatAdministrators() *BotCommandScopeAllChatAdministrators {
	if val, ok := bcs.impl.(*BotCommandScopeAllChatAdministrators); ok {
		return val
	}
	return nil
}

func (bcs *BotCommandScopeData) Chat() *BotCommandScopeChat {
	if val, ok := bcs.impl.(*BotCommandScopeChat); ok {
		return val
	}
	return nil
}

func (bcs *BotCommandScopeData) ChatMember() *BotCommandScopeChatMember {
	if val, ok := bcs.impl.(*BotCommandScopeChatMember); ok {
		return val
	}
	return nil
}

// BotCommandScopeDefault represents the default scope of bot commands. Default
// commands are used if no commands with a narrower scope are specified for the
// user.
//
// See https://core.telegram.org/bots/api#botcommandscopedefault
type BotCommandScopeDefault struct {
	Type string `json:"type"` // always "default"
}

func (bcs *BotCommandScopeDefault) isBotCommandScope() {}

// BotCommandScopeAllPrivateChats represents the scope of bot commands, covering
// all private chats.
//
// See https://core.telegram.org/bots/api#botcommandscopeallprivatechats
type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"` // always "all_private_chats"
}

func (bcs *BotCommandScopeAllPrivateChats) isBotCommandScope() {}

// BotCommandScopeAllGroupChats represents the scope of bot commands, covering
// all group and supergroup chats.
//
// See https://core.telegram.org/bots/api#botcommandscopeallgroupchats
type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"` // always "all_group_chats"
}

func (bcs *BotCommandScopeAllGroupChats) isBotCommandScope() {}

// BotCommandScopeAllChatAdministrators represents the scope of bot commands,
// covering all group and supergroup chat administrators.
//
// See https://core.telegram.org/bots/api#botcommandscopeallchatadministrators
type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"` // always "all_chat_administrators"
}

func (bcs *BotCommandScopeAllChatAdministrators) isBotCommandScope() {}

// BotCommandScopeChat represents the scope of bot commands, covering a specific chat.
//
// See https://core.telegram.org/bots/api#botcommandscopechat
type BotCommandScopeChat struct {
	Type   string `json:"type"` // always "chat"
	ChatID any    `json:"chat_id"`
}

func (bcs *BotCommandScopeChat) isBotCommandScope() {}

// BotCommandScopeChatAdministrators represents the scope of bot commands,
// covering all administrators of a specific group or supergroup chat.
//
// See https://core.telegram.org/bots/api#botcommandscopechatadministrators
type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"` // always "chat_administrators"
	ChatID any    `json:"chat_id"`
}

func (bcs *BotCommandScopeChatAdministrators) isBotCommandScope() {}

// BotCommandScopeChatMember represents the scope of bot commands, covering a
// specific member of a group or supergroup chat.
//
// See https://core.telegram.org/bots/api#botcommandscopechatmember
type BotCommandScopeChatMember struct {
	Type   string `json:"type"` // always "chat_member"
	ChatID any    `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

func (bcs *BotCommandScopeChatMember) isBotCommandScope() {}
