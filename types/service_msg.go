package types

// ProximityAlertTriggered represents the content of a service message, sent
// whenever a user in the chat triggers a proximity alert set by another user.
//
// See https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

// MessageAutoDeleteTimerChanged represents a service message about a change in
// auto-delete timer settings.
//
// See https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// WriteAccessAllowed represents a service message about a user allowing a bot to
// write messages after adding it to the attachment menu, launching a Web App
// from a link, or accepting an explicit request from a Web App sent by the
// method RequestWriteAccess.
//
// See https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

// PaidMessagePriceChanged describes a service message about a change in the
// price of paid messages within a chat.
//
// See https://core.telegram.org/bots/api#paidmessagepricechanged
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}

// DirectMessagePriceChanged describes a service message about a change in the
// price of direct messages sent to a channel chat.
//
// See https://core.telegram.org/bots/api#directmessagepricechanged
type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   int  `json:"direct_message_star_count,omitempty"`
}
