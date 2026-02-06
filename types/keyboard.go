package types

// ReplyMarkup is an interface representing different types of reply markups
// that can be attached to messages in Telegram. Current implementations include
// ReplyKeyboardMarkup, ReplyKeyboardRemove, InlineKeyboardMarkup, and ForceReply.
type ReplyMarkup interface {
	isReplyMarkup()
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options (see
// Introduction to bots for details and examples). Not supported in channels and
// for messages sent on behalf of a Telegram Business account.
//
// See https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          bool               `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
	Selective             bool               `json:"selective,omitempty"`
}

func (rkm *ReplyKeyboardMarkup) isReplyMarkup() {}

// ReplyKeyboardRemove upon receiving a message with this object, Telegram
// clients will remove the current custom keyboard and display the default
// letter-keyboard. By default, custom keyboards are displayed until a new
// keyboard is sent by a bot. An exception is made for one-time keyboards that
// are hidden immediately after the user presses a button (see
// ReplyKeyboardMarkup). Not supported in channels and for messages sent on
// behalf of a Telegram Business account.
//
// See https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

func (rkr *ReplyKeyboardRemove) isReplyMarkup() {}

// KeyboardButton represents one button of the reply keyboard. At most one of the
// optional fields must be used to specify type of the button. For simple text
// buttons, String can be used instead of this object to specify the button text.
//
// See https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            string                      `json:"text"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  bool                        `json:"request_contact,omitempty"`
	RequestLocation bool                        `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
}

// KeyboardButtonRequestUsers defines the criteria used to request suitable
// users. Information about the selected users will be shared with the bot when
// the corresponding button is pressed.
//
// See https://core.telegram.org/bots/api#keyboardbuttonrequestusers
type KeyboardButtonRequestUsers struct {
	RequestID       int32 `json:"request_id"`
	UserIsBot       bool  `json:"user_is_bot,omitempty"`
	UserIsPremium   bool  `json:"user_is_premium,omitempty"`
	MaxQuantity     int   `json:"max_quantity,omitempty"`
	RequestName     bool  `json:"request_name,omitempty"`
	RequestUsername bool  `json:"request_username,omitempty"`
	RequestPhoto    bool  `json:"request_photo,omitempty"`
}

// KeyboardButtonRequestChat defines the criteria used to request a suitable
// chat. Information about the selected chat will be shared with the bot when the
// corresponding button is pressed. The bot will be granted requested rights in
// the chat if appropriate.
//
// See https://core.telegram.org/bots/api#keyboardbuttonrequestchat
type KeyboardButtonRequestChat struct {
	RequestID               int32                    `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel"`
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`
	RequestTitle            bool                     `json:"request_title,omitempty"`
	RequestUsername         bool                     `json:"request_username,omitempty"`
	RequestPhoto            bool                     `json:"request_photo,omitempty"`
}

// KeyboardButtonPollType represents type of  poll, which is allowed to be created
// and sent when the corresponding button is pressed.
//
// See https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to
// the message it belongs to.
//
// See https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (ik *InlineKeyboardMarkup) isReplyMarkup() {}

// InlineKeyboardButton represents one button of an inline keyboard. Exactly one
// of the optional fields must be used to specify type of the button.
//
// See https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	URL                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginURL                     *LoginURL                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}

// LoginURL represents a parameter of the inline keyboard button used to
// automatically authorize a user. Serves as a great replacement for the Telegram
// Login Widget when the user is coming from Telegram.
//
// See https://core.telegram.org/bots/api#loginurl
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// SwitchInlineQueryChosenChat represents an inline button that switches the
// current user to inline mode in a chosen chat, with an optional default inline
// query.
//
// See https://core.telegram.org/bots/api#switchinlinequerychosenchat
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

// CopyTextButton represents an inline keyboard button that copies specified text
// to the clipboard.
//
// See https://core.telegram.org/bots/api#copytextbutton
type CopyTextButton struct {
	Text string `json:"text"`
}

// ForceReply upon receiving a message with this object, Telegram clients will
// display a reply interface to the user (act as if the user has selected the
// bot's message and tapped 'Reply'). This can be extremely useful if you want to
// create user-friendly step-by-step interfaces without having to sacrifice
// privacy mode. Not supported in channels and for messages sent on behalf of a
// Telegram Business account.
//
// See https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

func (fr *ForceReply) isReplyMarkup() {}
