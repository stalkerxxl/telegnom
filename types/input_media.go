package types

import (
	"encoding/json"
	"fmt"
)

// InputMedia represents the content of a media message to be sent. It should be one of
// InputMediaAnimation || InputMediaDocument || InputMediaAudio || InputMediaPhoto || InputMediaVideo
//
// See https://core.telegram.org/bots/api#inputmedia
type InputMedia interface {
	isInputMedia()
}

// InputMediaData is a wrapper for InputMedia interface.
//
// See https://core.telegram.org/bots/api#inputmedia
type InputMediaData struct {
	impl InputMedia
}

func (im *InputMediaData) MarshalJSON() ([]byte, error) {
	if im.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(im.impl)
}

func (im *InputMediaData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "video":
		var val InputMediaVideo
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		im.impl = &val
	case "photo":
		var val InputMediaPhoto
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		im.impl = &val
	case "audio":
		var val InputMediaAudio
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		im.impl = &val
	case "document":
		var val InputMediaDocument
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		im.impl = &val
	case "image":
		var val InputMediaAnimation
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		im.impl = &val
	default:
		return fmt.Errorf("unknown InputMedia type: %s", helper.Type)
	}

	return nil
}

func (im *InputMediaData) Animation() *InputMediaAnimation {
	if v, ok := im.impl.(*InputMediaAnimation); ok {
		return v
	}
	return nil
}

func (im *InputMediaData) Document() *InputMediaDocument {
	if v, ok := im.impl.(*InputMediaDocument); ok {
		return v
	}
	return nil
}

func (im *InputMediaData) Audio() *InputMediaAudio {
	if v, ok := im.impl.(*InputMediaAudio); ok {
		return v
	}
	return nil
}

func (im *InputMediaData) Photo() *InputMediaPhoto {
	if v, ok := im.impl.(*InputMediaPhoto); ok {
		return v
	}
	return nil
}

func (im *InputMediaData) Video() *InputMediaVideo {
	if v, ok := im.impl.(*InputMediaVideo); ok {
		return v
	}
	return nil
}

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC
// video without sound) to be sent.
//
// See https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type                  string          `json:"type"` // must be "animation"
	Media                 *InputFile      `json:"media"`
	Thumbnail             *InputFile      `json:"thumbnail,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	Duration              int             `json:"duration,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

func (ima *InputMediaAnimation) isInputMedia() {}

// InputMediaDocument represents a general file to be sent.
//
// See https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type                        string          `json:"type"` // must be "document"
	Media                       *InputFile      `json:"media"`
	Thumbnail                   *InputFile      `json:"thumbnail,omitempty"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
}

func (imd *InputMediaDocument) isInputMedia() {}

// InputMediaAudio represents an audio file to be treated as music to be sent.
//
// See https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type            string          `json:"type"` // must be "audio"
	Media           *InputFile      `json:"media"`
	Thumbnail       *InputFile      `json:"thumbnail,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        int             `json:"duration,omitempty"`
	Performer       string          `json:"performer,omitempty"`
	Title           string          `json:"title,omitempty"`
}

func (ima *InputMediaAudio) isInputMedia() {}

// InputMediaPhoto represents a photo to be sent.
//
// See https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type                  string          `json:"type"` // must be "photo"
	Media                 *InputFile      `json:"media"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

func (imp *InputMediaPhoto) isInputMedia() {}

// InputMediaVideo represents a video to be sent.
//
// See https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type                  string          `json:"type"` // must be "video"
	Media                 *InputFile      `json:"media"`
	Thumbnail             *InputFile      `json:"thumbnail,omitempty"`
	Cover                 *InputFile      `json:"cover,omitempty"`
	StartTimestamp        int             `json:"start_timestamp,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	Duration              int             `json:"duration,omitempty"`
	SupportsStreaming     bool            `json:"supports_streaming,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

func (imv *InputMediaVideo) isInputMedia() {}
