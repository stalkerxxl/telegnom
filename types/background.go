package types

import (
	"encoding/json"
	"fmt"
)

// BackgroundFill describes the way a background is filled based on the selected
// colors. Currently, it can be one of [BackgroundFillSolid] ||
// [BackgroundFillGradient] || [BackgroundFillFreeformGradient].
//
// See https://core.telegram.org/bots/api#backgroundfill
type BackgroundFill interface {
	isBackgroundFill()
}

// BackgroundFillData is a wrapper for the BackgroundFill interface.
//
// See https://core.telegram.org/bots/api#backgroundfill
type BackgroundFillData struct {
	impl BackgroundFill
}

func (bfd *BackgroundFillData) MarshalJSON() ([]byte, error) {
	if bfd.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(bfd.impl)
}

func (bfd *BackgroundFillData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}
	switch helper.Type {
	case "solid":
		var val BackgroundFillSolid
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bfd.impl = &val
	case "gradient":
		var val BackgroundFillGradient
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bfd.impl = &val
	case "freeform_gradient":
		var val BackgroundFillFreeformGradient
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bfd.impl = &val
	default:
		return fmt.Errorf("unknown BackgroundFillData type: %s", helper.Type)
	}
	return nil
}

func (bfd *BackgroundFillData) Solid() *BackgroundFillSolid {
	if val, ok := bfd.impl.(*BackgroundFillSolid); ok {
		return val
	}
	return nil
}

func (bfd *BackgroundFillData) Gradient() *BackgroundFillGradient {
	if val, ok := bfd.impl.(*BackgroundFillGradient); ok {
		return val
	}
	return nil
}

func (bfd *BackgroundFillData) FreeformGradient() *BackgroundFillFreeformGradient {
	if val, ok := bfd.impl.(*BackgroundFillFreeformGradient); ok {
		return val
	}
	return nil
}

// BackgroundFillSolid is filled using the selected color.
// See https://core.telegram.org/bots/api#backgroundfillsolid
type BackgroundFillSolid struct {
	Type  string `json:"type"` // always "solid"
	Color int    `json:"color"`
}

func (bf *BackgroundFillSolid) isBackgroundFill() {}

// BackgroundFillGradient is a gradient fill.
// See https://core.telegram.org/bots/api#backgroundfillgradient
type BackgroundFillGradient struct {
	Type          string `json:"type"` // always "gradient"
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
}

func (bf *BackgroundFillGradient) isBackgroundFill() {}

// BackgroundFillFreeformGradient is a freeform gradient that rotates after every message in the chat.
// See https://core.telegram.org/bots/api#backgroundfillfreeformgradient
type BackgroundFillFreeformGradient struct {
	Type   string `json:"type"` // always "freeform_gradient"
	Colors []int  `json:"colors"`
}

func (bf *BackgroundFillFreeformGradient) isBackgroundFill() {}

// BackgroundType describes the type of background. Currently, it can be one of
// BackgroundTypeFill || BackgroundTypeWallpaper || BackgroundTypePattern || BackgroundTypeChatTheme.
//
// See https://core.telegram.org/bots/api#backgroundtype
type BackgroundType interface {
	isBackgroundType()
}

// BackgroundTypeData is a wrapper for the BackgroundType interface.
//
// See https://core.telegram.org/bots/api#backgroundtype
type BackgroundTypeData struct {
	impl BackgroundType
}

func (bt *BackgroundTypeData) MarshalJSON() ([]byte, error) {
	if bt.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(bt.impl)
}

func (bt *BackgroundTypeData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}
	switch helper.Type {
	case "fill":
		var val BackgroundTypeFill
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bt.impl = &val
	case "wallpaper":
		var val BackgroundTypeWallpaper
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bt.impl = &val
	case "pattern":
		var val BackgroundTypePattern
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bt.impl = &val
	case "chat_theme":
		var val BackgroundTypeChatTheme
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		bt.impl = &val
	default:
		return fmt.Errorf("unknown BackgroundType type: %s", helper.Type)
	}
	return nil
}

func (bt *BackgroundTypeData) Fill() *BackgroundTypeFill {
	if val, ok := bt.impl.(*BackgroundTypeFill); ok {
		return val
	}
	return nil
}

func (bt *BackgroundTypeData) Wallpaper() *BackgroundTypeWallpaper {
	if val, ok := bt.impl.(*BackgroundTypeWallpaper); ok {
		return val
	}
	return nil
}

func (bt *BackgroundTypeData) Pattern() *BackgroundTypePattern {
	if val, ok := bt.impl.(*BackgroundTypePattern); ok {
		return val
	}
	return nil
}

func (bt *BackgroundTypeData) ChatTheme() *BackgroundTypeChatTheme {
	if val, ok := bt.impl.(*BackgroundTypeChatTheme); ok {
		return val
	}
	return nil
}

// BackgroundTypeFill is automatically filled based on the selected colors.
// See https://core.telegram.org/bots/api#backgroundtypefill
type BackgroundTypeFill struct {
	Type             string              `json:"type"` // always "fill"
	Fill             *BackgroundFillData `json:"fill"`
	DarkThemeDimming int                 `json:"dark_theme_dimming"`
}

func (bt *BackgroundTypeFill) isBackgroundType() {}

// BackgroundTypeWallpaper is a wallpaper in the JPEG format.
// See https://core.telegram.org/bots/api#backgroundtypewallpaper
type BackgroundTypeWallpaper struct {
	Type             string    `json:"type"` // always "wallpaper"
	Document         *Document `json:"document"`
	DarkThemeDimming int       `json:"dark_theme_dimming"`
	IsBlurred        bool      `json:"is_blurred,omitempty"`
	IsMoving         bool      `json:"is_moving,omitempty"`
}

func (bt *BackgroundTypeWallpaper) isBackgroundType() {}

// BackgroundTypePattern is a .PNG or .TGV
// (gzipped subset of SVG with MIME type “application/x-tgwallpattern”)
// pattern to be combined with the background fill chosen by the user.
// See https://core.telegram.org/bots/api#backgroundtypepattern
type BackgroundTypePattern struct {
	Type       string              `json:"type"` // always "pattern"
	Document   *Document           `json:"document"`
	Fill       *BackgroundFillData `json:"fill"`
	Intensity  int                 `json:"intensity"`
	IsInverted bool                `json:"is_inverted,omitempty"`
	IsMoving   bool                `json:"is_moving,omitempty"`
}

func (bt *BackgroundTypePattern) isBackgroundType() {}

// BackgroundTypeChatTheme is taken directly from a built-in chat theme.
// See https://core.telegram.org/bots/api#backgroundtypechattheme
type BackgroundTypeChatTheme struct {
	Type      string `json:"type"` // always "chat_theme"
	ThemeName string `json:"theme_name"`
}

func (bt *BackgroundTypeChatTheme) isBackgroundType() {}

// ChatBackground represents a chat background.
// See https://core.telegram.org/bots/api#chatbackground
type ChatBackground struct {
	Type *BackgroundTypeData `json:"type"`
}
