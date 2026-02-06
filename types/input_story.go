package types

import (
	"encoding/json"
	"fmt"
)

// InputStoryContent is an interface representing the content of a story to be
// posted. Currently, it can be either a InputStoryContentPhoto or
// InputStoryContentVideo.
//
// See https://core.telegram.org/bots/api#inputstorycontent
type InputStoryContent interface {
	isInputStoryContent()
}

// InputStoryContentData is a wrapper for the InputStoryContent interface.
//
// See https://core.telegram.org/bots/api#inputstorycontent
type InputStoryContentData struct {
	impl InputStoryContent
}

func (isc *InputStoryContentData) MarshalJSON() ([]byte, error) {
	if isc.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(isc.impl)
}

func (isc *InputStoryContentData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "photo":
		var val InputStoryContentPhoto
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		isc.impl = &val
	case "video":
		var val InputStoryContentVideo
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		isc.impl = &val
	default:
		return fmt.Errorf("unknown InputStoryContent type: %s", helper.Type)
	}

	return nil
}

func (isc *InputStoryContentData) Photo() *InputStoryContentPhoto {
	if v, ok := isc.impl.(*InputStoryContentPhoto); ok {
		return v
	}
	return nil
}

func (isc *InputStoryContentData) Video() *InputStoryContentVideo {
	if v, ok := isc.impl.(*InputStoryContentVideo); ok {
		return v
	}
	return nil
}

// InputStoryContentPhoto describes a photo to post as a story.
//
// See https://core.telegram.org/bots/api#inputstorycontentphoto
type InputStoryContentPhoto struct {
	Type  string     `json:"type"` // must be "photo"
	Photo *InputFile `json:"photo"`
}

func (isc *InputStoryContentPhoto) isInputStoryContent() {}

// InputStoryContentVideo describes a video to post as a story.
//
// See https://core.telegram.org/bots/api#inputstorycontentvideo
type InputStoryContentVideo struct {
	Type                string     `json:"type"` // must be "video"
	Video               *InputFile `json:"video"`
	Duration            float64    `json:"duration,omitempty"`
	CoverFrameTimestamp float64    `json:"cover_frame_timestamp,omitempty"`
	IsAnimation         bool       `json:"is_animation,omitempty"`
}

func (isc *InputStoryContentVideo) isInputStoryContent() {}
