package types

import (
	"encoding/json"
	"fmt"
)

// InputProfilePhoto is an interface representing a profile photo to set.
// Currently, it can be one of InputProfilePhotoStatic || InputProfilePhotoAnimated
//
// See https://core.telegram.org/bots/api#inputprofilephoto
type InputProfilePhoto interface {
	isInputProfilePhoto()
}

// InputProfilePhotoData is a wrapper struct for the InputProfilePhoto interface.
//
// See https://core.telegram.org/bots/api#inputprofilephoto
type InputProfilePhotoData struct {
	impl InputProfilePhoto
}

func (ipp *InputProfilePhotoData) MarshalJSON() ([]byte, error) {
	if ipp.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(ipp.impl)
}

func (ipp *InputProfilePhotoData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "static":
		var val InputProfilePhotoStatic
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		ipp.impl = &val
	case "animated":
		var val InputProfilePhotoAnimated
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		ipp.impl = &val
	default:
		return fmt.Errorf("unknown InputProfilePhoto type: %s", helper.Type)
	}

	return nil
}

func (ipp *InputProfilePhotoData) Static() *InputProfilePhotoStatic {
	if val, ok := ipp.impl.(*InputProfilePhotoStatic); ok {
		return val
	}
	return nil
}

func (ipp *InputProfilePhotoData) Animated() *InputProfilePhotoAnimated {
	if val, ok := ipp.impl.(*InputProfilePhotoAnimated); ok {
		return val
	}
	return nil
}

// InputProfilePhotoStatic - a static profile photo in the .JPG format.
//
// See https://core.telegram.org/bots/api#inputprofilephotostatic
type InputProfilePhotoStatic struct {
	Type  string     `json:"type"` // must be "static"
	Photo *InputFile `json:"photo"`
}

func (ipp *InputProfilePhotoStatic) isInputProfilePhoto() {}

// InputProfilePhotoAnimated - an animated profile photo in the MPEG4 format.
//
// See https://core.telegram.org/bots/api#inputprofilephotoanimated
type InputProfilePhotoAnimated struct {
	Type               string     `json:"type"` // must be "animated"
	Animation          *InputFile `json:"animation"`
	MainFrameTimestamp float64    `json:"main_frame_timestamp,omitempty"`
}

func (ipp *InputProfilePhotoAnimated) isInputProfilePhoto() {}
