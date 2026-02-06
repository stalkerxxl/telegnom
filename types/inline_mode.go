package types

// InlineQuery represents an incoming inline query. When the user sends an empty
// query, your bot could return some default or trending results.
//
// See https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType ChatType  `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// EffectiveUser returns the user who sent the inline query.
func (q *InlineQuery) EffectiveUser() *User {
	return q.From
}

// InlineQueryResultsButton represents a button to be shown above inline query
// results.
//
//	You must use exactly one of the optional fields.
//
// See https://core.telegram.org/bots/api#inlinequeryresultsbutton
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
	StartParameter string      `json:"start_parameter,omitempty"`
}

// ChosenInlineResult represents a InlineQueryResult of an inline query that was chosen by
// the user and sent to their chat partner.
//
// See https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// EffectiveUser returns the user who chose the result.
func (r *ChosenInlineResult) EffectiveUser() *User {
	return r.From
}

// PreparedInlineMessage describes an inline message to be sent by a user of a Mini App.
//
// See https://core.telegram.org/bots/api#preparedinlinemessage
type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int    `json:"expiration_date"`
}
