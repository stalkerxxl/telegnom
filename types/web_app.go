package types

// WebAppData describes data sent from a Web App to the bot.
//
// See https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// WebAppInfo describes a Web App.
//
// See https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	URL string `json:"url"`
}

// SentWebAppMessage describes an inline message sent by a Web App on behalf of a user.
//
// See https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}
