package types

import (
	"encoding/json"
	"fmt"
)

// ChatBoostSource describes the source of a chat boost. It can be one of
// ChatBoostSourcePremium || ChatBoostSourceGiftCode || ChatBoostSourceGiveaway
//
// See https://core.telegram.org/bots/api#chatboostsource
type ChatBoostSource interface {
	isChatBoostSource()
}

// ChatBoostSourceData is a wrapper around ChatBoostSource
//
// See https://core.telegram.org/bots/api#chatboostsource
type ChatBoostSourceData struct {
	impl ChatBoostSource
}

func (cbs *ChatBoostSourceData) MarshalJSON() ([]byte, error) {
	if cbs.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(cbs.impl)
}

func (cbs *ChatBoostSourceData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Source {
	case "premium":
		var val ChatBoostSourcePremium
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cbs.impl = &val
	case "gift_code":
		var val ChatBoostSourceGiftCode
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cbs.impl = &val
	case "giveaway":
		var val ChatBoostSourceGiveaway
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cbs.impl = &val
	default:
		return fmt.Errorf("unknown ChatBoostSource type: %s", helper.Source)
	}

	return nil
}

func (cbs *ChatBoostSourceData) Premium() *ChatBoostSourcePremium {
	if v, ok := cbs.impl.(*ChatBoostSourcePremium); ok {
		return v
	}
	return nil
}

func (cbs *ChatBoostSourceData) GiftCode() *ChatBoostSourceGiftCode {
	if v, ok := cbs.impl.(*ChatBoostSourceGiftCode); ok {
		return v
	}
	return nil
}

func (cbs *ChatBoostSourceData) Giveaway() *ChatBoostSourceGiveaway {
	if v, ok := cbs.impl.(*ChatBoostSourceGiveaway); ok {
		return v
	}
	return nil
}

// ChatBoostSourcePremium was obtained by subscribing to Telegram Premium
// or by gifting a Telegram Premium subscription to another user.
//
// See https://core.telegram.org/bots/api#chatboostsourcepremium
type ChatBoostSourcePremium struct {
	Source string `json:"source"` // always “premium”
	User   *User  `json:"user"`
}

func (cbs *ChatBoostSourcePremium) isChatBoostSource() {}

// ChatBoostSourceGiftCode was obtained by the creation of Telegram Premium
// gift codes to boost a chat. Each such code boosts the chat 4 times
// for the duration of the corresponding Telegram Premium subscription.
//
// See https://core.telegram.org/bots/api#chatboostsourcegiftcode
type ChatBoostSourceGiftCode struct {
	Source string `json:"source"` // always “gift_code”
	User   *User  `json:"user"`
}

func (cbs *ChatBoostSourceGiftCode) isChatBoostSource() {}

// ChatBoostSourceGiveaway was obtained by the creation of a Telegram Premium or a Telegram Star giveaway.
// This boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription
// for Telegram Premium giveaways and prize_star_count / 500 times for one year for Telegram Star giveaways.
//
// See https://core.telegram.org/bots/api#chatboostsourcegiveaway
type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"` // always “giveaway”
	GiveawayMessageID int    `json:"giveaway_message_id"`
	User              *User  `json:"user,omitempty"`
	PrizeStarCount    int    `json:"prize_star_count,omitempty"`
	IsUnclaimed       bool   `json:"is_unclaimed,omitempty"`
}

func (cbs *ChatBoostSourceGiveaway) isChatBoostSource() {}

// ChatBoost contains information about a chat boost.
//
// See https://core.telegram.org/bots/api#chatboost
type ChatBoost struct {
	BoostID        string               `json:"boost_id"`
	AddDate        int                  `json:"add_date"`
	ExpirationDate int                  `json:"expiration_date"`
	Source         *ChatBoostSourceData `json:"source"`
}

// ChatBoostUpdated represents a boost added to a chat or changed.
//
// See https://core.telegram.org/bots/api#chatboostupdated
type ChatBoostUpdated struct {
	Chat  *Chat      `json:"chat"`
	Boost *ChatBoost `json:"boost"`
}

// ChatBoostRemoved represents a boost removed from a chat.
//
// See https://core.telegram.org/bots/api#chatboostremoved
type ChatBoostRemoved struct {
	Chat       *Chat                `json:"chat"`
	BoostID    string               `json:"boost_id"`
	RemoveDate int                  `json:"remove_date"`
	Source     *ChatBoostSourceData `json:"source"`
}

// UserChatBoosts represents a list of boosts added to a chat by a user.
//
// See https://core.telegram.org/bots/api#userchatboosts
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}
