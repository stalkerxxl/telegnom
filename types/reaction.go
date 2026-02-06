package types

import (
	"encoding/json"
	"fmt"
)

// ReactionType is an interface for different reaction types. Currently, it can
// be one of ReactionTypeEmoji || ReactionTypeCustomEmoji || ReactionTypePaid
//
// See https://core.telegram.org/bots/api#reactiontype
type ReactionType interface {
	isReactionType()
}

// ReactionTypeData is a wrapper for the ReactionType interface.
//
// See https://core.telegram.org/bots/api#reactiontype
type ReactionTypeData struct {
	impl ReactionType
}

func (rt *ReactionTypeData) MarshalJSON() ([]byte, error) {
	if rt.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(rt.impl)
}

func (rt *ReactionTypeData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}
	switch helper.Type {
	case "emoji":
		var val ReactionTypeEmoji
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		rt.impl = &val
	case "custom_emoji":
		var val ReactionTypeCustomEmoji
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		rt.impl = &val
	case "paid":
		var val ReactionTypePaid
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		rt.impl = &val
	default:
		return fmt.Errorf("unknown ReactionType %s", helper.Type)
	}

	return nil
}

func (rt *ReactionTypeData) Emoji() *ReactionTypeEmoji {
	if v, ok := rt.impl.(*ReactionTypeEmoji); ok {
		return v
	}
	return nil
}

func (rt *ReactionTypeData) CustomEmoji() *ReactionTypeCustomEmoji {
	if v, ok := rt.impl.(*ReactionTypeCustomEmoji); ok {
		return v
	}
	return nil
}

func (rt *ReactionTypeData) Paid() *ReactionTypePaid {
	if v, ok := rt.impl.(*ReactionTypePaid); ok {
		return v
	}
	return nil
}

// ReactionTypeEmoji is based on an emoji.
//
// See https://core.telegram.org/bots/api#reactiontypeemoji
type ReactionTypeEmoji struct {
	Type  string `json:"type"` // must be "emoji"
	Emoji string `json:"emoji"`
}

func (r *ReactionTypeEmoji) isReactionType() {}

// ReactionTypeCustomEmoji is based on a custom emoji.
//
// See https://core.telegram.org/bots/api#reactiontypecustomemoji
type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"` // must be "custom_emoji"
	CustomEmojiID string `json:"custom_emoji_id"`
}

func (r *ReactionTypeCustomEmoji) isReactionType() {}

// ReactionTypePaid is a paid reaction.
//
// See https://core.telegram.org/bots/api#reactiontypepaid
type ReactionTypePaid struct {
	Type string `json:"type"` // must be "paid"
}

func (r *ReactionTypePaid) isReactionType() {}

// ReactionCount represents a reaction added to a message along with the number
// of times it was added.
//
// See https://core.telegram.org/bots/api#reactioncount
type ReactionCount struct {
	Type       *ReactionTypeData `json:"type"`
	TotalCount int               `json:"total_count"`
}

// MessageReactionUpdated represents a change of a reaction on a message
// performed by a user.
//
// See https://core.telegram.org/bots/api#messagereactionupdated
type MessageReactionUpdated struct {
	Chat        *Chat              `json:"chat"`
	MessageID   int                `json:"message_id"`
	User        *User              `json:"user,omitempty"`
	ActorChat   *Chat              `json:"actor_chat,omitempty"`
	Date        int                `json:"date"`
	OldReaction []ReactionTypeData `json:"old_reaction"`
	NewReaction []ReactionTypeData `json:"new_reaction"`
}

// MessageReactionCountUpdated represents reaction changes on a message with
// anonymous reactions.
//
// See https://core.telegram.org/bots/api#messagereactioncountupdated
type MessageReactionCountUpdated struct {
	Chat      *Chat           `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}
