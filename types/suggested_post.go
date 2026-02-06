package types

// SuggestedPostApproved describes a service message about the approval of a
// suggested post.
//
// See https://core.telegram.org/bots/api#suggestedpostapproved
type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
	SendDate             int                 `json:"send_date,omitempty"`
}

// SuggestedPostApprovalFailed describes a service message about the failed
// approval of a suggested post. Currently, only caused by insufficient user
// funds at the time of approval.
//
// See https://core.telegram.org/bots/api#suggestedpostapprovalfailed
type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
}

// SuggestedPostDeclined describes a service message about the rejection of a
// suggested post.
//
// See https://core.telegram.org/bots/api#suggestedpostdeclined
type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Comment              string   `json:"comment,omitempty"`
}

// SuggestedPostPaid describes a service message about a successful payment for a
// suggested post.
//
// See https://core.telegram.org/bots/api#suggestedpostpaid
type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message,omitempty"`
	Currency             string      `json:"currency"`
	Amount               int         `json:"amount,omitempty"`
	StarAmount           *StarAmount `json:"star_amount,omitempty"`
}

// SuggestedPostRefunded describes a service message about a payment refund for a
// suggested post.
//
// See https://core.telegram.org/bots/api#suggestedpostrefunded
type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Reason               string   `json:"reason"`
}

// SuggestedPostPrice describes the price of a suggested post.
//
// See https://core.telegram.org/bots/api#suggestedpostprice
type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

// SuggestedPostInfo contains information about a suggested post.
//
// See https://core.telegram.org/bots/api#suggestedpostinfo
type SuggestedPostInfo struct {
	State    string              `json:"state"` // “pending”, “approved”, “declined”
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int                 `json:"send_date,omitempty"`
}

// SuggestedPostParameters contains parameters of a post that is being suggested
// by the bot.
//
// See https://core.telegram.org/bots/api#suggestedpostparameters
type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int                 `json:"send_date,omitempty"`
}
