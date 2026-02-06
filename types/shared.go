package types

// SharedUser contains information about a user that was shared with the bot
// using a KeyboardButtonRequestUsers button.
//
// See https://core.telegram.org/bots/api#shareduser
type SharedUser struct {
	UserID    int64      `json:"user_id"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Username  string     `json:"username,omitempty"`
	Photo     PhotoSizes `json:"photo,omitempty"`
}

// UsersShared contains information about the users whose identifiers were shared
// with the bot using a KeyboardButtonRequestUsers button.
//
// See https://core.telegram.org/bots/api#usershared
type UsersShared struct {
	RequestID int          `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

// ChatShared contains information about a chat that was shared with the bot
// using a KeyboardButtonRequestChat button.
//
// See https://core.telegram.org/bots/api#chatshared
type ChatShared struct {
	RequestID int        `json:"request_id"`
	ChatID    int64      `json:"chat_id"`
	Title     string     `json:"title,omitempty"`
	Username  string     `json:"username,omitempty"`
	Photo     PhotoSizes `json:"photo,omitempty"`
}
