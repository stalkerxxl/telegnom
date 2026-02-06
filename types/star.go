package types

// StarAmount describes an amount of Telegram Stars.
//
// See https://core.telegram.org/bots/api#staramount
type StarAmount struct {
	Amount         int `json:"amount"`
	NanostarAmount int `json:"nanostar_amount,omitempty"`
}
