package types

import (
	"encoding/json"
	"fmt"
)

// MessageOrBool represents a union type that can either be a Message object or a boolean value.
//
// See https://core.telegram.org/bots/api#updating-messages
type MessageOrBool struct {
	msg *Message
	ok  *bool
}

func (mb *MessageOrBool) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("empty byte slice")
	}

	switch b[0] {
	case 't', 'f': // JSON true or false
		var val bool
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		mb.ok = &val
	case '{': // JSON object
		var msg Message
		if err := json.Unmarshal(b, &msg); err != nil {
			return err
		}
		mb.msg = &msg
	default:
		return fmt.Errorf("unexpected json type: %s", string(b))
	}

	return nil
}

func (mb *MessageOrBool) IsMessage() bool {
	return mb.msg != nil
}

// Message returns the Message object, if it exists, otherwise nil
func (mb *MessageOrBool) Message() *Message {
	return mb.msg
}

// Success returns the boolean value, if it exists, otherwise false
func (mb *MessageOrBool) Success() bool {
	if mb.ok != nil {
		return *mb.ok
	}
	return false
}
