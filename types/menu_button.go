package types

import (
	"encoding/json"
	"fmt"
)

// MenuButton is an interface representing different types of menu buttons. It
// should be one of MenuButtonCommands || MenuButtonWebApp || MenuButtonDefault
//
// See https://core.telegram.org/bots/api#menubutton
type MenuButton interface {
	isMenuButton()
}

// MenuButtonData is a wrapper struct for the MenuButton interface.
//
// See https://core.telegram.org/bots/api#menubutton
type MenuButtonData struct {
	impl MenuButton
}

func (mb *MenuButtonData) MarshalJSON() ([]byte, error) {
	if mb.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(mb.impl)
}

func (mb *MenuButtonData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "commands":
		var val MenuButtonCommands
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		mb.impl = &val
	case "web_app":
		var val MenuButtonWebApp
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		mb.impl = &val
	case "default":
		var val MenuButtonDefault
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		mb.impl = &val
	default:
		return fmt.Errorf("unknown MenuButton type: %s", helper.Type)
	}

	return nil
}

func (mb *MenuButtonData) Commands() *MenuButtonCommands {
	if val, ok := mb.impl.(*MenuButtonCommands); ok {
		return val
	}
	return nil
}

func (mb *MenuButtonData) WebApp() *MenuButtonWebApp {
	if val, ok := mb.impl.(*MenuButtonWebApp); ok {
		return val
	}
	return nil
}

func (mb *MenuButtonData) Default() *MenuButtonDefault {
	if val, ok := mb.impl.(*MenuButtonDefault); ok {
		return val
	}
	return nil
}

// MenuButtonCommands represents a menu button, which opens the bot's list of
// commands.
//
// See https://core.telegram.org/bots/api#menubuttoncommands
type MenuButtonCommands struct {
	Type string `json:"type"` // must be "commands"
}

func (mb *MenuButtonCommands) isMenuButton() {}

// MenuButtonWebApp represents a menu button, which launches a WebApp.
//
// See https://core.telegram.org/bots/api#menubuttonwebapp
type MenuButtonWebApp struct {
	Type   string      `json:"type"` // must be "web_app"
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app"`
}

func (mb *MenuButtonWebApp) isMenuButton() {}

// MenuButtonDefault describes that no specific value for the menu button was
// set.
//
// See https://core.telegram.org/bots/api#menubuttondefault
type MenuButtonDefault struct {
	Type string `json:"type"` // must be "default"
}

func (mb *MenuButtonDefault) isMenuButton() {}
