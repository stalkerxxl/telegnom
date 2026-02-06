package types

import (
	"encoding/json"
	"fmt"
)

// Story represents a story.
//
// See https://core.telegram.org/bots/api#story
type Story struct {
	Chat *Chat `json:"chat"`
	ID   int   `json:"id"`
}

// StoryArea describes a clickable area on a story media.
//
// See https://core.telegram.org/bots/api#storyarea
type StoryArea struct {
	Position *StoryAreaPosition `json:"position"`
	Type     *StoryAreaTypeData `json:"type"`
}

// StoryAreaPosition describes the position of a clickable area within a story.
//
// See https://core.telegram.org/bots/api#storyareaposition
type StoryAreaPosition struct {
	XPercentage            float64 `json:"x_percentage"`
	YPercentage            float64 `json:"y_percentage"`
	WidthPercentage        float64 `json:"width_percentage"`
	HeightPercentage       float64 `json:"height_percentage"`
	RotationAngle          float64 `json:"rotation_angle"`
	CornerRadiusPercentage float64 `json:"corner_radius_percentage"`
}

// StoryAreaType is an interface for different types of StoryAreaType. Currently,
// it can be one of StoryAreaTypeLocation || StoryAreaTypeSuggestedReaction ||
// StoryAreaTypeLink || StoryAreaTypeWeather || StoryAreaTypeUniqueGift
//
// See https://core.telegram.org/bots/api#storyareatype
type StoryAreaType interface {
	isStoryAreaType()
}

// StoryAreaTypeData is a wrapper for the different types of StoryAreaType.
//
// See https://core.telegram.org/bots/api#storyareatype
type StoryAreaTypeData struct {
	impl StoryAreaType
}

func (sat *StoryAreaTypeData) MarshalJSON() ([]byte, error) {
	if sat.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(sat.impl)
}

func (sat *StoryAreaTypeData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "location":
		var val StoryAreaTypeLocation
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		sat.impl = &val
	case "suggested_reaction":
		var val StoryAreaTypeSuggestedReaction
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		sat.impl = &val
	case "link":
		var val StoryAreaTypeLink
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		sat.impl = &val
	case "weather":
		var val StoryAreaTypeWeather
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		sat.impl = &val
	case "unique_gift":
		var val StoryAreaTypeUniqueGift
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		sat.impl = &val
	default:
		return fmt.Errorf("unknown StoryAreaType type: %sat", helper.Type)
	}

	return nil
}

func (sat *StoryAreaTypeData) Location() *StoryAreaTypeLocation {
	if v, ok := sat.impl.(*StoryAreaTypeLocation); ok {
		return v
	}
	return nil
}

func (sat *StoryAreaTypeData) SuggestedReaction() *StoryAreaTypeSuggestedReaction {
	if v, ok := sat.impl.(*StoryAreaTypeSuggestedReaction); ok {
		return v
	}
	return nil
}

func (sat *StoryAreaTypeData) Link() *StoryAreaTypeLink {
	if v, ok := sat.impl.(*StoryAreaTypeLink); ok {
		return v
	}
	return nil
}

func (sat *StoryAreaTypeData) Weather() *StoryAreaTypeWeather {
	if v, ok := sat.impl.(*StoryAreaTypeWeather); ok {
		return v
	}
	return nil
}

func (sat *StoryAreaTypeData) UniqueGift() *StoryAreaTypeUniqueGift {
	if v, ok := sat.impl.(*StoryAreaTypeUniqueGift); ok {
		return v
	}
	return nil
}

// StoryAreaTypeLocation describes a story area pointing to a location.
// Currently, a story can have up to 10 location areas.
//
// See https://core.telegram.org/bots/api#storyareatypelocation
type StoryAreaTypeLocation struct {
	Type      string           `json:"type"` // always "location"
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
	Address   *LocationAddress `json:"address,omitempty"`
}

func (sat *StoryAreaTypeLocation) isStoryAreaType() {}

// StoryAreaTypeSuggestedReaction describes a story area pointing to a suggested
// reaction. Currently, a story can have up to 5 suggested reaction areas.
//
// See https://core.telegram.org/bots/api#storyareatypesuggestedreaction
type StoryAreaTypeSuggestedReaction struct {
	Type         string            `json:"type"` // always "suggested_reaction"
	ReactionType *ReactionTypeData `json:"reaction_type"`
	IsDark       bool              `json:"is_dark,omitempty"`
	IsFlipped    bool              `json:"is_flipped,omitempty"`
}

func (sat *StoryAreaTypeSuggestedReaction) isStoryAreaType() {}

// StoryAreaTypeLink describes a story area pointing to an HTTP or types:// link.
// Currently, a story can have up to 3 link areas.
//
// See https://core.telegram.org/bots/api#storyareatypelink
type StoryAreaTypeLink struct {
	Type string `json:"type"` // always "link"
	URL  string `json:"url"`
}

func (sat *StoryAreaTypeLink) isStoryAreaType() {}

// StoryAreaTypeWeather describes a story area containing weather information.
// Currently, a story can have up to 3 weather areas.
//
// See https://core.telegram.org/bots/api#storyareatypeweather
type StoryAreaTypeWeather struct {
	Type            string  `json:"type"` // always "weather"
	Temperature     float64 `json:"temperature"`
	Emoji           string  `json:"emoji"`
	BackgroundColor int     `json:"background_color"`
}

func (sat *StoryAreaTypeWeather) isStoryAreaType() {}

// StoryAreaTypeUniqueGift describes a story area pointing to a unique gift.
// Currently, a story can have at most 1 unique gift area.
//
// See https://core.telegram.org/bots/api#storyareatypeuniquegift
type StoryAreaTypeUniqueGift struct {
	Type string `json:"type"` // always "unique_gift"
	Name string `json:"name"`
}

func (sat *StoryAreaTypeUniqueGift) isStoryAreaType() {}
