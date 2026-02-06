package types

import (
	"encoding/json"
	"fmt"
)

// GiftBackground describes the background of a gift.
//
// See https://core.telegram.org/bots/api#giftbackground
type GiftBackground struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	TextColor   int `json:"text_color"`
}

// Gift represents a gift that can be sent by the bot.
//
// See https://core.telegram.org/bots/api#gift
type Gift struct {
	ID                     string          `json:"id"`
	Sticker                *Sticker        `json:"sticker"`
	StarCount              int             `json:"star_count"`
	UpgradeStarCount       int             `json:"upgrade_star_count,omitempty"`
	IsPremium              bool            `json:"is_premium,omitempty"`
	HasColors              bool            `json:"has_colors,omitempty"`
	TotalCount             int             `json:"total_count,omitempty"`
	RemainingCount         int             `json:"remaining_count,omitempty"`
	PersonalTotalCount     int             `json:"personal_total_count,omitempty"`
	PersonalRemainingCount int             `json:"personal_remaining_count,omitempty"`
	Background             *GiftBackground `json:"background,omitempty"`
	UniqueGiftVariantCount int             `json:"unique_gift_variant_count,omitempty"`
	PublisherChat          *Chat           `json:"publisher_chat,omitempty"`
}

// Gifts represent a list of gifts.
//
// See https://core.telegram.org/bots/api#gifts
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

// UniqueGiftModel describes the model of a unique gift.
//
// See https://core.telegram.org/bots/api#uniquegiftmodel
type UniqueGiftModel struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int      `json:"rarity_per_mille"`
}

// UniqueGiftSymbol describes the symbol shown on the pattern of a unique gift.
//
// See https://core.telegram.org/bots/api#uniquegiftsymbol
type UniqueGiftSymbol struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int      `json:"rarity_per_mille"`
}

// UniqueGiftBackdropColors describes the colors of the backdrop of a unique gift.
//
// See https://core.telegram.org/bots/api#uniquegiftbackdropcolors
type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

// UniqueGiftBackdrop describes the backdrop of a unique gift.
//
// See https://core.telegram.org/bots/api#uniquegiftbackdrop
type UniqueGiftBackdrop struct {
	Name           string                    `json:"name"`
	Colors         *UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                       `json:"rarity_per_mille"`
}

// UniqueGiftColors contains information about the color scheme for a user's
// name, message replies and link previews based on a unique gift.
//
// See https://core.telegram.org/bots/api#uniquegiftcolors
type UniqueGiftColors struct {
	ModelCustomEmojiID   string `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID  string `json:"symbol_custom_emoji_id"`
	LightThemeMainColor  int    `json:"light_theme_main_color"`
	LightThemeOtherColor []int  `json:"light_theme_other_color"`
	DarkThemeMainColor   int    `json:"dark_theme_main_color"`
	DarkThemeOtherColor  []int  `json:"dark_theme_other_color"`
}

// UniqueGift describes a unique gift that was upgraded from a regular gift.
//
// See https://core.telegram.org/bots/api#uniquegift
type UniqueGift struct {
	GiftID           string              `json:"gift_id"`
	BaseName         string              `json:"base_name"`
	Name             string              `json:"name"`
	Number           int                 `json:"number"`
	Model            *UniqueGiftModel    `json:"model"`
	Symbol           *UniqueGiftSymbol   `json:"symbol"`
	Backdrop         *UniqueGiftBackdrop `json:"backdrop"`
	IsPremium        bool                `json:"is_premium,omitempty"`
	IsFromBlockchain bool                `json:"is_from_blockchain,omitempty"`
	Colors           *UniqueGiftColors   `json:"colors,omitempty"`
	PublisherChat    *Chat               `json:"publisher_chat,omitempty"`
}

// GiftInfo describes a service message about a regular gift that was sent or received.
//
// See https://core.telegram.org/bots/api#giftinfo
type GiftInfo struct {
	Gift                    *Gift           `json:"gift"`
	OwnedGiftID             string          `json:"owned_gift_id,omitempty"`
	ConvertStarCount        int             `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int             `json:"prepaid_upgrade_star_count,omitempty"`
	IsUpgradeSeparate       bool            `json:"is_upgrade_separate,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
	UniqueGiftNumber        int             `json:"unique_gift_number,omitempty"`
}

// UniqueGiftInfo describes a service message about a unique gift that was sent or received.
//
// See https://core.telegram.org/bots/api#uniquegiftinfo
type UniqueGiftInfo struct {
	Gift                *UniqueGift `json:"gift"`
	Origin              string      `json:"origin"`
	LastResaleStarCount int         `json:"last_resale_star_count,omitempty"`
	LastResaleAmount    int         `json:"last_resale_amount,omitempty"`
	OwnedGiftID         string      `json:"owned_gift_id,omitempty"`
	TransferStarCount   int         `json:"transfer_star_count,omitempty"`
	NextTransferDate    int         `json:"next_transfer_date,omitempty"`
}

// OwnedGift is an interface implemented by all types that describe owned gifts.
// Currently, it can be one of OwnedGiftRegular || OwnedGiftUnique
//
// See https://core.telegram.org/bots/api#ownedgift
type OwnedGift interface {
	isOwnedGift()
}

// OwnedGiftData is wrapper over OwnedGift interface.
//
// See https://core.telegram.org/bots/api#ownedgift
type OwnedGiftData struct {
	impl OwnedGift
}

func (og *OwnedGiftData) Regular() *OwnedGiftRegular {
	if val, ok := og.impl.(*OwnedGiftRegular); ok {
		return val
	}
	return nil
}

func (og *OwnedGiftData) Unique() *OwnedGiftUnique {
	if val, ok := og.impl.(*OwnedGiftUnique); ok {
		return val
	}
	return nil
}

func (og *OwnedGiftData) MarshalJSON() ([]byte, error) {
	if og.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(og.impl)
}

func (og *OwnedGiftData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "regular":
		var val OwnedGiftRegular
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		og.impl = &val
	case "unique":
		var val OwnedGiftUnique
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		og.impl = &val
	default:
		return fmt.Errorf("unknown OwnedGiftData type: %s", helper.Type)
	}

	return nil
}

// OwnedGiftRegular describes a regular gift owned by a user or a chat.
//
// See https://core.telegram.org/bots/api#ownedgiftregular
type OwnedGiftRegular struct {
	Type                    string          `json:"type"` // always "regular"
	Gift                    *Gift           `json:"gift"`
	OwnedGiftID             string          `json:"owned_gift_id,omitempty"`
	SenderUser              *User           `json:"sender_user,omitempty"`
	SendDate                int             `json:"send_date"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
	IsSaved                 bool            `json:"is_saved,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	WasRefunded             bool            `json:"was_refunded,omitempty"`
	ConvertStarCount        int             `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int             `json:"prepaid_upgrade_star_count,omitempty"`
	IsUpgradeSeparate       bool            `json:"is_upgrade_separate,omitempty"`
	UniqueGiftNumber        int             `json:"unique_gift_number,omitempty"`
}

func (og *OwnedGiftRegular) isOwnedGift() {}

// OwnedGiftUnique describes a unique gift received and owned by a user or a chat.
//
// See https://core.telegram.org/bots/api#ownedgiftunique
type OwnedGiftUnique struct {
	Type              string      `json:"type"` // always "unique"
	Gift              *UniqueGift `json:"gift"`
	OwnedGiftID       string      `json:"owned_gift_id,omitempty"`
	SenderUser        *User       `json:"sender_user,omitempty"`
	SendDate          int         `json:"send_date"`
	IsSaved           bool        `json:"is_saved,omitempty"`
	CanBeTransferred  bool        `json:"can_be_transferred,omitempty"`
	TransferStarCount int         `json:"transfer_star_count,omitempty"`
	NextTransferDate  int         `json:"next_transfer_date,omitempty"`
}

func (og *OwnedGiftUnique) isOwnedGift() {}

// OwnedGifts contains the list of gifts received and owned by a user or a chat.
//
// See https://core.telegram.org/bots/api#ownedgifts
type OwnedGifts struct {
	TotalCount int             `json:"total_count"`
	Gifts      []OwnedGiftData `json:"gifts"`
	NextOffset string          `json:"next_offset,omitempty"`
}

// AcceptedGiftTypes describes the types of gifts that can be gifted to a user or a chat.
//
// See https://core.telegram.org/bots/api#acceptedgifttypes
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription,omitempty"`
	GiftsFromChannels   bool `json:"gifts_from_channels,omitempty"`
}
