package types

import (
	"encoding/json"
	"fmt"
)

// InputPaidMedia is an interface representing the paid media to be sent.
// It can be one of InputPaidMediaPhoto or InputPaidMediaVideo.
//
// See https://core.telegram.org/bots/api#inputpaidmedia
type InputPaidMedia interface {
	isInputPaidMedia()
}

// InputPaidMediaData is a wrapper struct for the InputPaidMedia interface.
//
// See https://core.telegram.org/bots/api#inputpaidmedia
type InputPaidMediaData struct {
	impl InputPaidMedia
}

func (ipm *InputPaidMediaData) MarshalJSON() ([]byte, error) {
	if ipm.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(ipm.impl)
}

func (ipm *InputPaidMediaData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "photo":
		var val InputPaidMediaPhoto
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		ipm.impl = &val
	case "video":
		var val InputPaidMediaVideo
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		ipm.impl = &val
	default:
		return fmt.Errorf("unknown InputPaidMedia type: %s", helper.Type)
	}
	return nil
}

func (ipm *InputPaidMediaData) Photo() *InputPaidMediaPhoto {
	if v, ok := ipm.impl.(*InputPaidMediaPhoto); ok {
		return v
	}
	return nil
}

func (ipm *InputPaidMediaData) Video() *InputPaidMediaVideo {
	if v, ok := ipm.impl.(*InputPaidMediaVideo); ok {
		return v
	}
	return nil
}

// InputPaidMediaPhoto - the paid media to send is a photo.
//
// See https://core.telegram.org/bots/api#inputpaidmediaphoto
type InputPaidMediaPhoto struct {
	Type  string     `json:"type"` // must be "photo"
	Media *InputFile `json:"media"`
}

func (ipm *InputPaidMediaPhoto) isInputPaidMedia() {}

// InputPaidMediaVideo - the paid media to send is a video.
//
// See https://core.telegram.org/bots/api#inputpaidmediavideo
type InputPaidMediaVideo struct {
	Type              string     `json:"type"` // must be "video"
	Media             *InputFile `json:"media"`
	Thumbnail         *InputFile `json:"thumbnail,omitempty"`
	Cover             *InputFile `json:"cover,omitempty"`
	StartTimestamp    int        `json:"start_timestamp,omitempty"`
	Width             int        `json:"width,omitempty"`
	Height            int        `json:"height,omitempty"`
	Duration          int        `json:"duration,omitempty"`
	SupportsStreaming bool       `json:"supports_streaming,omitempty"`
}

func (ipm *InputPaidMediaVideo) isInputPaidMedia() {}
