package types

import (
	"encoding/json"
	"fmt"
)

// MessageOrigin is interface for different types of message origins. Can be one
// of the following types: MessageOriginUser || MessageOriginHiddenUser ||
// MessageOriginChat || MessageOriginChannel
//
// See https://core.telegram.org/bots/api#messageorigin
type MessageOrigin interface {
	isMessageOrigin()
}

// MessageOriginData is an interface wrapper for the different types of
// MessageOrigin.
//
// See https://core.telegram.org/bots/api#messageorigin
type MessageOriginData struct {
	impl MessageOrigin
}

func (mo *MessageOriginData) MarshalJSON() ([]byte, error) {
	if mo.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(mo.impl)
}

func (mo *MessageOriginData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "user":
		var val MessageOriginUser
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		mo.impl = &val
	case "hidden_user":
		var val MessageOriginHiddenUser
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		mo.impl = &val
	case "chat":
		var val MessageOriginChat
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		mo.impl = &val
	case "channel":
		var val MessageOriginChannel
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		mo.impl = &val
	default:
		return fmt.Errorf("unknown MessageOrigin type: %s", helper.Type)
	}

	return nil
}

func (mo *MessageOriginData) User() *MessageOriginUser {
	if val, ok := mo.impl.(*MessageOriginUser); ok {
		return val
	}
	return nil
}

func (mo *MessageOriginData) HiddenUser() *MessageOriginHiddenUser {
	if val, ok := mo.impl.(*MessageOriginHiddenUser); ok {
		return val
	}
	return nil
}

func (mo *MessageOriginData) Chat() *MessageOriginChat {
	if val, ok := mo.impl.(*MessageOriginChat); ok {
		return val
	}
	return nil
}

func (mo *MessageOriginData) Channel() *MessageOriginChannel {
	if val, ok := mo.impl.(*MessageOriginChannel); ok {
		return val
	}
	return nil
}

// MessageOriginUser - the message was originally sent by a known user.
//
// See https://core.telegram.org/bots/api#messageoriginuser
type MessageOriginUser struct {
	Type       string `json:"type"` // always “user”
	Date       int    `json:"date"`
	SenderUser *User  `json:"sender_user"`
}

func (mo *MessageOriginUser) isMessageOrigin() {}

// MessageOriginHiddenUser - the message was originally sent by an unknown user.
//
// See https://core.telegram.org/bots/api#messageoriginhiddenuser
type MessageOriginHiddenUser struct {
	Type           string `json:"type"` // always “hidden_user”
	Date           int    `json:"date"`
	SenderUserName string `json:"sender_user_name"`
}

func (mo *MessageOriginHiddenUser) isMessageOrigin() {}

// MessageOriginChat - the message was originally sent on behalf of a chat to a
// group chat.
//
// See https://core.telegram.org/bots/api#messageoriginchat
type MessageOriginChat struct {
	Type            string `json:"type"` // always “chat”
	Date            int    `json:"date"`
	SenderChat      *Chat  `json:"sender_chat"`
	AuthorSignature string `json:"author_signature,omitempty"`
}

func (mo *MessageOriginChat) isMessageOrigin() {}

// MessageOriginChannel - the message was originally sent to a channel chat.
//
// See https://core.telegram.org/bots/api#messageoriginchannel
type MessageOriginChannel struct {
	Type            string `json:"type"` // always “channel”
	Date            int    `json:"date"`
	Chat            *Chat  `json:"chat"`
	MessageID       int    `json:"message_id"`
	AuthorSignature string `json:"author_signature,omitempty"`
}

func (mo *MessageOriginChannel) isMessageOrigin() {}
