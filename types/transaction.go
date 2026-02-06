package types

import (
	"encoding/json"
	"fmt"
)

// TransactionPartner is an interface for different types of transaction partners.
// Currently, the following types are supported: TransactionPartnerUser, TransactionPartnerChat,
// TransactionPartnerAffiliateProgram, TransactionPartnerFragment, TransactionPartnerTelegramAds,
// TransactionPartnerTelegramApi, TransactionPartnerOther.
//
// See https://core.telegram.org/bots/api#transactionpartner
type TransactionPartner interface {
	isTransactionPartner()
}

// TransactionPartnerData is a wrapper for the TransactionPartner.
//
// See https://core.telegram.org/bots/api#transactionpartner
type TransactionPartnerData struct {
	impl TransactionPartner
}

func (tp *TransactionPartnerData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case "user":
		var val TransactionPartnerUser
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		tp.impl = &val
	case "chat":
		var val TransactionPartnerChat
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		tp.impl = &val
	case "affiliate_program":
		var val TransactionPartnerAffiliateProgram
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		tp.impl = &val
	case "fragment":
		var val TransactionPartnerFragment
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		tp.impl = &val
	case "telegram_ads":
		var val TransactionPartnerTelegramAds
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		tp.impl = &val
	case "telegram_api":
		var val TransactionPartnerTelegramApi
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		tp.impl = &val
	case "other":
		var val TransactionPartnerOther
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		tp.impl = &val
	default:
		return fmt.Errorf("unsupported TransactionPartnerData type: %s", helper.Type)
	}
	return nil
}

func (tp *TransactionPartnerData) User() *TransactionPartnerUser {
	if v, ok := tp.impl.(*TransactionPartnerUser); ok {
		return v
	}
	return nil
}

func (tp *TransactionPartnerData) Chat() *TransactionPartnerChat {
	if v, ok := tp.impl.(*TransactionPartnerChat); ok {
		return v
	}
	return nil
}

func (tp *TransactionPartnerData) AffiliateProgram() *TransactionPartnerAffiliateProgram {
	if v, ok := tp.impl.(*TransactionPartnerAffiliateProgram); ok {
		return v
	}
	return nil
}

func (tp *TransactionPartnerData) Fragment() *TransactionPartnerFragment {
	if v, ok := tp.impl.(*TransactionPartnerFragment); ok {
		return v
	}
	return nil
}

func (tp *TransactionPartnerData) TelegramAds() *TransactionPartnerTelegramAds {
	if v, ok := tp.impl.(*TransactionPartnerTelegramAds); ok {
		return v
	}
	return nil
}

func (tp *TransactionPartnerData) TelegramApi() *TransactionPartnerTelegramApi {
	if v, ok := tp.impl.(*TransactionPartnerTelegramApi); ok {
		return v
	}
	return nil
}

func (tp *TransactionPartnerData) Other() *TransactionPartnerOther {
	if v, ok := tp.impl.(*TransactionPartnerOther); ok {
		return v
	}
	return nil
}

// TransactionPartnerUser describes a transaction with a user.
//
// See https://core.telegram.org/bots/api#transactionpartneruser
type TransactionPartnerUser struct {
	Type                        string          `json:"type"` // must be "user"
	TransactionType             string          `json:"transaction_type"`
	User                        *User           `json:"user"`
	Affiliate                   *AffiliateInfo  `json:"affiliate,omitempty"`
	InvoicePayload              string          `json:"invoice_payload,omitempty"`
	SubscriptionPeriod          int             `json:"subscription_period,omitempty"`
	PaidMedia                   []PaidMediaData `json:"paid_media,omitempty"`
	PaidMediaPayload            string          `json:"paid_media_payload,omitempty"`
	Gift                        string          `json:"gift,omitempty"`
	PremiumSubscriptionDuration int             `json:"premium_subscription_duration,omitempty"`
}

func (tp *TransactionPartnerUser) isTransactionPartner() {}

// TransactionPartnerChat describes a transaction with a chat.
//
// See https://core.telegram.org/bots/api#transactionpartnerchat
type TransactionPartnerChat struct {
	Type string `json:"type"` // must be "chat"
	Chat *Chat  `json:"chat"`
	Gift *Gift  `json:"gift,omitempty"`
}

func (tp *TransactionPartnerChat) isTransactionPartner() {}

// TransactionPartnerAffiliateProgram describes the affiliate program that issued
// the affiliate commission received via this transaction.
//
// See https://core.telegram.org/bots/api#transactionpartneraffiliateprogram
type TransactionPartnerAffiliateProgram struct {
	Type               string `json:"type"` // must be "affiliate_program"
	SponsorUser        *User  `json:"sponsor_user,omitempty"`
	CommissionPerMille int    `json:"commission_per_mille"`
}

func (tp *TransactionPartnerAffiliateProgram) isTransactionPartner() {}

// TransactionPartnerFragment describes a withdrawal transaction with Fragment.
//
// See https://core.telegram.org/bots/api#transactionpartnerfragment
type TransactionPartnerFragment struct {
	Type            string                      `json:"type"` // must be "fragment"
	WithdrawalState *RevenueWithdrawalStateData `json:"withdrawal_state,omitempty"`
}

func (tp *TransactionPartnerFragment) isTransactionPartner() {}

// TransactionPartnerTelegramAds describes a withdrawal transaction to the
// Telegram Ads platform.
//
// See https://core.telegram.org/bots/api#transactionpartnertelegramads
type TransactionPartnerTelegramAds struct {
	Type string `json:"type"` // must be "telegram_ads"
}

func (tp *TransactionPartnerTelegramAds) isTransactionPartner() {}

// TransactionPartnerTelegramApi describes a transaction with payment for paid
// broadcasting.
//
// See https://core.telegram.org/bots/api#transactionpartnertelegramapi
type TransactionPartnerTelegramApi struct {
	Type         string `json:"type"` // must be "telegram_api"
	RequestCount int    `json:"request_count"`
}

func (tp *TransactionPartnerTelegramApi) isTransactionPartner() {}

// TransactionPartnerOther describes a transaction with an unknown source or recipient.
//
// See https://core.telegram.org/bots/api#transactionpartnerother
type TransactionPartnerOther struct {
	Type string `json:"type"` // must be "other"
}

func (tp *TransactionPartnerOther) isTransactionPartner() {}

// StarTransaction describes a Telegram Star transaction. Note that if the buyer
// initiates a chargeback with the payment provider from whom they acquired Stars
// (e.g., Apple, Google) following this transaction, the refunded Stars will be
// deducted from the bot's balance. This is outside of Telegram's control.
//
// See https://core.telegram.org/bots/api#startransaction
type StarTransaction struct {
	ID             string                  `json:"id"`
	Amount         int                     `json:"amount"`
	NanostarAmount int                     `json:"nanostar_amount,omitempty"`
	Date           int                     `json:"date"`
	Source         *TransactionPartnerData `json:"source,omitempty"`
	Receiver       *TransactionPartnerData `json:"receiver,omitempty"`
}

// StarTransactions contains a list of Telegram Star transactions.
//
// See https://core.telegram.org/bots/api#startransactions
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}
