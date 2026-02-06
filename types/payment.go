package types

import (
	"encoding/json"
	"fmt"
)

// LabeledPrice represents a portion of the price for goods or services.
//
// [LabeledPrice.Amount] of the product in the smallest units of the currency
// (integer, not float/double). For example, for a price of US$ 1.45 pass amount
// = 145. See the exp parameter in currencies.json, it shows the number of digits
// past the decimal point for each currency (2 for the majority of currencies).
//
// See https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// Invoice contains basic information about an invoice.
//
// See https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// ShippingAddress represents a shipping address.
//
// See https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// OrderInfo represents information about an order.
//
// See https://core.telegram.org/bots/api#orderinfo
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingOption represents one shipping option.
//
// See https://core.telegram.org/bots/api#shippingoption
type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

// SuccessfulPayment contains basic information about a successful payment. Note
// that if the buyer initiates a chargeback with the relevant payment provider
// following this transaction, the funds may be debited from your balance. This
// is outside of Telegram's control.
//
// See https://core.telegram.org/bots/api#successfulpayment
type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`
	TotalAmount                int        `json:"total_amount"`
	InvoicePayload             string     `json:"invoice_payload"`
	SubscriptionExpirationDate int        `json:"subscription_expiration_date,omitempty"`
	IsRecurring                bool       `json:"is_recurring,omitempty"`
	IsFirstRecurring           bool       `json:"is_first_recurring,omitempty"`
	ShippingOptionID           string     `json:"shipping_option_id,omitempty"`
	OrderInfo                  *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID    string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID    string     `json:"provider_payment_charge_id"`
}

// RefundedPayment contains basic information about a refunded payment.
//
// See https://core.telegram.org/bots/api#refundedpayment
type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int    `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}

// ShippingQuery contains information about an incoming shipping query.
//
// See https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// EffectiveUser returns the user who sent the shipping query.
func (q *ShippingQuery) EffectiveUser() *User {
	return q.From
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
//
// See https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

// EffectiveUser returns the user who sent the pre-checkout query.
func (q *PreCheckoutQuery) EffectiveUser() *User {
	return q.From
}

// PaidMediaPurchased contains information about a paid media purchase.
//
// See https://core.telegram.org/bots/api#paidmediapurchased
type PaidMediaPurchased struct {
	From             *User  `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

type RevenueWithdrawalStateType string

const (
	RevenueWithdrawalStateTypePending   RevenueWithdrawalStateType = "pending"
	RevenueWithdrawalStateTypeSucceeded RevenueWithdrawalStateType = "succeeded"
	RevenueWithdrawalStateTypeFailed    RevenueWithdrawalStateType = "failed"
)

// RevenueWithdrawalState is an interface representing the state of a revenue
// withdrawal. Can be one of the following types: RevenueWithdrawalStatePending
// || RevenueWithdrawalStateSucceeded || RevenueWithdrawalStateFailed
//
// See https://core.telegram.org/bots/api#revenuewithdrawalstate
type RevenueWithdrawalState interface {
	isRevenueWithdrawalState()
}

// RevenueWithdrawalStateData is a wrapper for the RevenueWithdrawalState.
//
// See https://core.telegram.org/bots/api#revenuewithdrawalstate
type RevenueWithdrawalStateData struct {
	impl RevenueWithdrawalState
}

func (rws *RevenueWithdrawalStateData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Type {
	case string(RevenueWithdrawalStateTypePending):
		var val RevenueWithdrawalStatePending
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		rws.impl = &val
	case string(RevenueWithdrawalStateTypeSucceeded):
		var val RevenueWithdrawalStateSucceeded
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		rws.impl = &val
	case string(RevenueWithdrawalStateTypeFailed):
		var val RevenueWithdrawalStateFailed
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		rws.impl = &val
	default:
		return fmt.Errorf("unsupported RevenueWithdrawalState type: %s", helper.Type)
	}

	return nil
}

func (rws *RevenueWithdrawalStateData) Pending() *RevenueWithdrawalStatePending {
	if v, ok := rws.impl.(*RevenueWithdrawalStatePending); ok {
		return v
	}
	return nil
}

func (rws *RevenueWithdrawalStateData) Succeeded() *RevenueWithdrawalStateSucceeded {
	if v, ok := rws.impl.(*RevenueWithdrawalStateSucceeded); ok {
		return v
	}
	return nil
}

func (rws *RevenueWithdrawalStateData) Failed() *RevenueWithdrawalStateFailed {
	if v, ok := rws.impl.(*RevenueWithdrawalStateFailed); ok {
		return v
	}
	return nil
}

// RevenueWithdrawalStatePending - the withdrawal is in progress.
//
// See https://core.telegram.org/bots/api#revenuewithdrawalstatepending
type RevenueWithdrawalStatePending struct {
	Type RevenueWithdrawalStateType `json:"type"`
}

func (rws *RevenueWithdrawalStatePending) isRevenueWithdrawalState() {}

// RevenueWithdrawalStateSucceeded - the withdrawal succeeded.
//
// See https://core.telegram.org/bots/api#revenuewithdrawalstatesucceeded
type RevenueWithdrawalStateSucceeded struct {
	Type RevenueWithdrawalStateType `json:"type"`
	Date int                        `json:"date"`
	URL  string                     `json:"url"`
}

func (rws *RevenueWithdrawalStateSucceeded) isRevenueWithdrawalState() {}

// RevenueWithdrawalStateFailed - the withdrawal failed and the transaction was
// refunded.
//
// See https://core.telegram.org/bots/api#revenuewithdrawalstatefailed
type RevenueWithdrawalStateFailed struct {
	Type RevenueWithdrawalStateType `json:"type"`
}

func (rws *RevenueWithdrawalStateFailed) isRevenueWithdrawalState() {}

// AffiliateInfo contains information about the affiliate that received a
// commission via this transaction.
//
// See https://core.telegram.org/bots/api#affiliateinfo
type AffiliateInfo struct {
	AffiliateUser      *User `json:"affiliate_user,omitempty"`
	AffiliateChat      *Chat `json:"affiliate_chat,omitempty"`
	CommissionPerMille int   `json:"commission_per_mille"`
	Amount             int   `json:"amount"`
	NanostarAmount     int   `json:"nanostar_amount,omitempty"`
}
