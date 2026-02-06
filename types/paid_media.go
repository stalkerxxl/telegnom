package types

import (
	"encoding/json"
	"fmt"
)

// PaidMediaInfo describes the paid media added to a message.
//
// See https://core.telegram.org/bots/api#paidmediainfo
type PaidMediaInfo struct {
	StarCount int             `json:"star_count"`
	PaidMedia []PaidMediaData `json:"paid_media"`
}

// PaidMedia is an interface for different types of paid media. Currently, it can
// be one of PaidMediaPreview || PaidMediaPhoto || PaidMediaVideo.
//
// See https://core.telegram.org/bots/api#paidmedia
type PaidMedia interface {
	isPaidMedia()
}

// PaidMediaData is a wrapper for the different types of paid media.
//
// See https://core.telegram.org/bots/api#paidmedia
type PaidMediaData struct {
	impl PaidMedia
}

func (pm *PaidMediaData) MarshalJSON() ([]byte, error) {
	if pm.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(pm.impl)
}

func (pm *PaidMediaData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "preview":
		var preview PaidMediaPreview
		if err := json.Unmarshal(b, &preview); err != nil {
			return err
		}
		pm.impl = &preview
	case "photo":
		var photo PaidMediaPhoto
		if err := json.Unmarshal(b, &photo); err != nil {
			return err
		}
		pm.impl = &photo
	case "video":
		var video PaidMediaVideo
		if err := json.Unmarshal(b, &video); err != nil {
			return err
		}
		pm.impl = &video
	default:
		return fmt.Errorf("unknown PaidMedia type: %s", helper.Type)
	}

	return nil
}

func (pm *PaidMediaData) Preview() *PaidMediaPreview {
	if v, ok := pm.impl.(*PaidMediaPreview); ok {
		return v
	}
	return nil
}

func (pm *PaidMediaData) Photo() *PaidMediaPhoto {
	if v, ok := pm.impl.(*PaidMediaPhoto); ok {
		return v
	}
	return nil
}

func (pm *PaidMediaData) Video() *PaidMediaVideo {
	if v, ok := pm.impl.(*PaidMediaVideo); ok {
		return v
	}
	return nil
}

// PaidMediaPreview - the paid media isn't available before the payment.
//
// See https://core.telegram.org/bots/api#paidmediapreview
type PaidMediaPreview struct {
	Type     string // must be "preview"
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

func (pm *PaidMediaPreview) isPaidMedia() {}

// PaidMediaPhoto - the paid media is a photo.
//
// See https://core.telegram.org/bots/api#paidmediaphoto
type PaidMediaPhoto struct {
	Type  string     // must be "photo"
	Photo PhotoSizes `json:"photo"`
}

func (pm *PaidMediaPhoto) isPaidMedia() {}

// PaidMediaVideo - the paid media is a video.
//
// See https://core.telegram.org/bots/api#paidmediavideo
type PaidMediaVideo struct {
	Type  string // must be "video"
	Video *Video `json:"video"`
}

func (pm *PaidMediaVideo) isPaidMedia() {}
