package bot

import (
	"github.com/stalkerxxl/telegnom/types"
)

// SendMessageParams https://core.telegram.org/bots/api#sendmessage
type SendMessageParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Text                    string                         `json:"text"`
	ParseMode               types.ParseMode                `json:"parse_mode,omitempty"`
	Entities                []types.MessageEntity          `json:"entities,omitempty"`
	LinkPreviewOptions      *types.LinkPreviewOptions      `json:"link_preview_options,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// ForwardMessageParams https://core.telegram.org/bots/api#forwardmessage
type ForwardMessageParams struct {
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	FromChatID              any                            `json:"from_chat_id"`
	VideoStartTimestamp     int                            `json:"video_start_timestamp,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	MessageID               int                            `json:"message_id"`
}

// ForwardMessagesParams https://core.telegram.org/bots/api#forwardmessages
type ForwardMessagesParams struct {
	ChatID                any   `json:"chat_id"`
	MessageThreadID       int   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID int   `json:"direct_messages_topic_id,omitempty"`
	FromChatID            any   `json:"from_chat_id"`
	MessageIDs            []int `json:"message_ids"`
	DisableNotification   bool  `json:"disable_notification,omitempty"`
	ProtectContent        bool  `json:"protect_content,omitempty"`
}

// CopyMessageParams https://core.telegram.org/bots/api#copymessage
type CopyMessageParams struct {
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	FromChatID              any                            `json:"from_chat_id"`
	MessageID               int                            `json:"message_id"`
	VideoStartTimestamp     int                            `json:"video_start_timestamp,omitempty"`
	Caption                 string                         `json:"caption,omitempty"`
	ParseMode               types.ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities         []types.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                           `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// CopyMessagesParams https://core.telegram.org/bots/api#copymessages
type CopyMessagesParams struct {
	ChatID                any   `json:"chat_id"`
	MessageThreadID       int   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID int   `json:"direct_messages_topic_id,omitempty"`
	FromChatID            any   `json:"from_chat_id"`
	MessageIDs            []int `json:"message_ids"`
	DisableNotification   bool  `json:"disable_notification,omitempty"`
	ProtectContent        bool  `json:"protect_content,omitempty"`
	RemoveCaption         bool  `json:"remove_caption,omitempty"`
}

// SendPhotoParams https://core.telegram.org/bots/api#sendphoto
type SendPhotoParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Photo                   *types.InputFile               `json:"photo,omitempty" media:"photo"`
	Caption                 string                         `json:"caption,omitempty"`
	ParseMode               types.ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities         []types.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                           `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                           `json:"has_spoiler,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendAudioParams https://core.telegram.org/bots/api#sendaudio
type SendAudioParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Audio                   *types.InputFile               `json:"audio,omitempty" media:"audio"`
	Caption                 string                         `json:"caption,omitempty"`
	ParseMode               types.ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities         []types.MessageEntity          `json:"caption_entities,omitempty"`
	Duration                int                            `json:"duration,omitempty"`
	Performer               string                         `json:"performer,omitempty"`
	Title                   string                         `json:"title,omitempty"`
	Thumbnail               *types.InputFile               `json:"thumbnail,omitempty" media:"thumbnail"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendDocumentParams https://core.telegram.org/bots/api#senddocument
type SendDocumentParams struct {
	BusinessConnectionID        string                         `json:"business_connection_id,omitempty"`
	ChatID                      any                            `json:"chat_id"`
	MessageThreadID             int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID       int                            `json:"direct_messages_topic_id,omitempty"`
	Document                    *types.InputFile               `json:"document,omitempty" media:"document"`
	Thumbnail                   *types.InputFile               `json:"thumbnail,omitempty" media:"thumbnail,omitempty"`
	Caption                     string                         `json:"caption,omitempty"`
	ParseMode                   types.ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities             []types.MessageEntity          `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                           `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                           `json:"disable_notification,omitempty"`
	ProtectContent              bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast          bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID             string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters     *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters             *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup                 types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendVideoParams https://core.telegram.org/bots/api#sendvideo
type SendVideoParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Video                   *types.InputFile               `json:"video,omitempty" media:"video"`
	Duration                int                            `json:"duration,omitempty"`
	Width                   int                            `json:"width,omitempty"`
	Height                  int                            `json:"height,omitempty"`
	Thumbnail               *types.InputFile               `json:"thumbnail,omitempty" media:"thumbnail"`
	Cover                   *types.InputFile               `json:"cover,omitempty" media:"cover"`
	StartTimestamp          int                            `json:"start_timestamp,omitempty"`
	Caption                 string                         `json:"caption,omitempty"`
	ParseMode               types.ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities         []types.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                           `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                           `json:"has_spoiler,omitempty"`
	SupportsStreaming       bool                           `json:"supports_streaming,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendAnimationParams https://core.telegram.org/bots/api#sendanimation
type SendAnimationParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Animation               *types.InputFile               `json:"animation,omitempty" media:"animation"`
	Duration                int                            `json:"duration,omitempty"`
	Width                   int                            `json:"width,omitempty"`
	Height                  int                            `json:"height,omitempty"`
	Thumbnail               *types.InputFile               `json:"thumbnail,omitempty" media:"thumbnail"`
	Caption                 string                         `json:"caption,omitempty"`
	ParseMode               types.ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities         []types.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                           `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                           `json:"has_spoiler,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendVoiceParams https://core.telegram.org/bots/api#sendvoice
type SendVoiceParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Voice                   *types.InputFile               `json:"voice,omitempty" media:"voice"`
	Caption                 string                         `json:"caption,omitempty"`
	ParseMode               types.ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities         []types.MessageEntity          `json:"caption_entities,omitempty"`
	Duration                int                            `json:"duration,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendVideoNoteParams https://core.telegram.org/bots/api#sendvideonote
type SendVideoNoteParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	VideoNote               *types.InputFile               `json:"video_note,omitempty" media:"video_note"`
	Duration                int                            `json:"duration,omitempty"`
	Length                  int                            `json:"length,omitempty"`
	Thumbnail               *types.InputFile               `json:"thumbnail,omitempty" media:"thumbnail,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendPaidMediaParams https://core.telegram.org/bots/api#sendpaidmedia
type SendPaidMediaParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	StarCount               int                            `json:"star_count"`
	Media                   []types.InputPaidMedia         `json:"media"`
	Payload                 string                         `json:"payload,omitempty"`
	Caption                 string                         `json:"caption,omitempty"`
	ParseMode               types.ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities         []types.MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                           `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

type SendMediaGroupParams struct {
	BusinessConnectionID string                 `json:"business_connection_id,omitempty"`
	ChatID               any                    `json:"chat_id"`
	MessageThreadID      int                    `json:"message_thread_id,omitempty"`
	Media                []types.InputMedia     `json:"media"`
	DisableNotification  bool                   `json:"disable_notification,omitempty"`
	ProtectContent       bool                   `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                   `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                 `json:"message_effect_id,omitempty"`
	ReplyParameters      *types.ReplyParameters `json:"reply_parameters,omitempty"`
}

// SendStoryParams https://core.telegram.org/bots/api#sendstory
type SendStoryParams struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"`
	ChatID               any                   `json:"chat_id"`
	Media                *types.InputFile      `json:"media,omitempty" media:"media"`
	Caption              string                `json:"caption,omitempty"`
	ParseMode            types.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities      []types.MessageEntity `json:"caption_entities,omitempty"`
	ReplyMarkup          types.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// SendLocationParams https://core.telegram.org/bots/api#sendlocation
type SendLocationParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Latitude                float64                        `json:"latitude"`
	Longitude               float64                        `json:"longitude"`
	HorizontalAccuracy      float64                        `json:"horizontal_accuracy,omitempty"`
	LivePeriod              int                            `json:"live_period,omitempty"`
	Heading                 int                            `json:"heading,omitempty"`
	ProximityAlertRadius    int                            `json:"proximity_alert_radius,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendVenueParams https://core.telegram.org/bots/api#sendvenue
type SendVenueParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Latitude                float64                        `json:"latitude"`
	Longitude               float64                        `json:"longitude"`
	Title                   string                         `json:"title"`
	Address                 string                         `json:"address"`
	FoursquareID            string                         `json:"foursquare_id,omitempty"`
	FoursquareType          string                         `json:"foursquare_type,omitempty"`
	GooglePlaceID           string                         `json:"google_place_id,omitempty"`
	GooglePlaceType         string                         `json:"google_place_type,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendContactParams https://core.telegram.org/bots/api#sendcontact
type SendContactParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	PhoneNumber             string                         `json:"phone_number"`
	FirstName               string                         `json:"first_name"`
	LastName                string                         `json:"last_name,omitempty"`
	VCard                   string                         `json:"vcard,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendPollParams https://core.telegram.org/bots/api#sendpoll
type SendPollParams struct {
	BusinessConnectionID  string                  `json:"business_connection_id,omitempty"`
	ChatID                any                     `json:"chat_id"`
	MessageThreadID       int                     `json:"message_thread_id,omitempty"`
	Question              string                  `json:"question"`
	QuestionParseMode     types.ParseMode         `json:"question_parse_mode,omitempty"`
	QuestionEntities      []types.MessageEntity   `json:"question_entities,omitempty"`
	Options               []types.InputPollOption `json:"options"`
	IsAnonymous           bool                    `json:"is_anonymous"`   // FIXME оставить коммент про разницу с дефолтом API
	Type                  string                  `json:"type,omitempty"` // FIXME сделать константы
	AllowsMultipleAnswers bool                    `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID       int                     `json:"correct_option_id,omitempty"`
	Explanation           string                  `json:"explanation,omitempty"`
	ExplanationParseMode  string                  `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities   []types.MessageEntity   `json:"explanation_entities,omitempty"`
	OpenPeriod            int                     `json:"open_period,omitempty"`
	CloseDate             int                     `json:"close_date,omitempty"`
	IsClosed              bool                    `json:"is_closed,omitempty"`
	DisableNotification   bool                    `json:"disable_notification,omitempty"`
	ProtectContent        bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID       string                  `json:"message_effect_id,omitempty"`
	ReplyParameters       *types.ReplyParameters  `json:"reply_parameters,omitempty"`
	ReplyMarkup           types.ReplyMarkup       `json:"reply_markup,omitempty"`
}

// SendChecklistParams https://core.telegram.org/bots/api#sendchecklist
type SendChecklistParams struct {
	BusinessConnectionID string                 `json:"business_connection_id"`
	ChatID               int                    `json:"chat_id"`
	Checklist            *types.InputChecklist  `json:"checklist"`
	DisableNotification  bool                   `json:"disable_notification,omitempty"`
	ProtectContent       bool                   `json:"protect_content,omitempty"`
	MessageEffectID      string                 `json:"message_effect_id,omitempty"`
	ReplyParameters      *types.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          types.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendDiceParams https://core.telegram.org/bots/api#senddice
type SendDiceParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Emoji                   string                         `json:"emoji,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendMessageDraftParams https://core.telegram.org/bots/api#sendmessagedraft
type SendMessageDraftParams struct {
	ChatID          int                   `json:"chat_id"`
	MessageThreadID int                   `json:"message_thread_id,omitempty"`
	DraftID         int                   `json:"draft_id"`
	Text            string                `json:"text"`
	ParseMode       types.ParseMode       `json:"parse_mode,omitempty"`
	Entities        []types.MessageEntity `json:"entities,omitempty"`
}

// SendChatActionParams https://core.telegram.org/bots/api#sendchataction
type SendChatActionParams struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageThreadID      int    `json:"message_thread_id,omitempty"`
	Action               string `json:"action"` // FIXME в константы!
}

// SetMessageReactionParams https://core.telegram.org/bots/api#setmessagereaction
type SetMessageReactionParams struct {
	ChatID    any                  `json:"chat_id"`
	MessageID int                  `json:"message_id"`
	Reaction  []types.ReactionType `json:"reaction,omitempty"`
	IsBig     bool                 `json:"is_big,omitempty"`
}

// GetUserProfilePhotosParams https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotosParams struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset,omitempty"`
	Limit  int   `json:"limit,omitempty"`
}

// SetUserEmojiStatusParams https://core.telegram.org/bots/api#setuseremojistatus
type SetUserEmojiStatusParams struct {
	UserID                    int64  `json:"user_id"`
	EmojiStatusCustomEmojiID  string `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate int    `json:"emoji_status_expiration_date,omitempty"`
}

// GetFileParams https://core.telegram.org/bots/api#getfile
type GetFileParams struct {
	FileID string `json:"file_id"`
}

// BanChatMemberParams https://core.telegram.org/bots/api#banchatmember
type BanChatMemberParams struct {
	ChatID         any   `json:"chat_id"`
	UserID         int64 `json:"user_id"`
	UntilDate      int   `json:"until_date,omitempty"`
	RevokeMessages bool  `json:"revoke_messages,omitempty"`
}

// UnbanChatMemberParams https://core.telegram.org/bots/api#unbanchatmember
type UnbanChatMemberParams struct {
	ChatID       any   `json:"chat_id"`
	UserID       int64 `json:"user_id"`
	OnlyIfBanned bool  `json:"only_if_banned,omitempty"`
}

// RestrictChatMemberParams https://core.telegram.org/bots/api#restrictchatmember
type RestrictChatMemberParams struct {
	ChatID                        any                    `json:"chat_id"`
	UserID                        int64                  `json:"user_id"`
	Permissions                   *types.ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool                   `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int                    `json:"until_date,omitempty"`
}

// PromoteChatMemberParams https://core.telegram.org/bots/api#promotechatmember
type PromoteChatMemberParams struct {
	ChatID                  any   `json:"chat_id"`
	UserID                  int64 `json:"user_id"`
	IsAnonymous             bool  `json:"is_anonymous,omitempty"`
	CanManageChat           bool  `json:"can_manage_chat,omitempty"`
	CanDeleteMessages       bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats     bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers      bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers       bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo           bool  `json:"can_change_info,omitempty"`
	CanInviteUsers          bool  `json:"can_invite_users,omitempty"`
	CanPostStories          bool  `json:"can_post_stories,omitempty"`
	CanEditStories          bool  `json:"can_edit_stories,omitempty"`
	CanDeleteStories        bool  `json:"can_delete_stories,omitempty"`
	CanPostMessages         bool  `json:"can_post_messages,omitempty"`
	CanEditMessages         bool  `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool  `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool  `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool  `json:"can_manage_direct_messages,omitempty"`
}

// SetChatAdministratorCustomTitleParams https://core.telegram.org/bots/api#setchatadministratorcustomtitle
type SetChatAdministratorCustomTitleParams struct {
	ChatID      any    `json:"chat_id"`
	UserID      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

// BanChatSenderChatParams https://core.telegram.org/bots/api#banchatsenderchat
type BanChatSenderChatParams struct {
	ChatID       any `json:"chat_id"`
	SenderChatID int `json:"sender_chat_id"`
}

// UnbanChatSenderChatParams https://core.telegram.org/bots/api#unbanchatsenderchat
type UnbanChatSenderChatParams struct {
	ChatID       any `json:"chat_id"`
	SenderChatID int `json:"sender_chat_id"`
}

// SetChatPermissionsParams https://core.telegram.org/bots/api#setchatpermissions
type SetChatPermissionsParams struct {
	ChatID                        any                    `json:"chat_id"`
	Permissions                   *types.ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool                   `json:"use_independent_chat_permissions,omitempty"`
}

// ExportChatInviteLinkParams https://core.telegram.org/bots/api#exportchatinvitelink
type ExportChatInviteLinkParams struct {
	ChatID any `json:"chat_id"`
}

// CreateChatInviteLinkParams https://core.telegram.org/bots/api#createchatinvitelink
type CreateChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

// EditChatInviteLinkParams https://core.telegram.org/bots/api#editchatinvitelink
type EditChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

// CreateChatSubscriptionInviteLinkParams https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
type CreateChatSubscriptionInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	SubscriptionPeriod int    `json:"subscription_period"`
	SubscriptionPrice  int    `json:"subscription_price"`
}

// EditChatSubscriptionInviteLinkParams https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
type EditChatSubscriptionInviteLinkParams struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
	Name       string `json:"name,omitempty"`
}

// RevokeChatInviteLinkParams https://core.telegram.org/bots/api#revokechatinvitelink
type RevokeChatInviteLinkParams struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

// ApproveChatJoinRequestParams https://core.telegram.org/bots/api#approvechatjoinrequest
type ApproveChatJoinRequestParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// DeclineChatJoinRequestParams https://core.telegram.org/bots/api#declinechatjoinrequest
type DeclineChatJoinRequestParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// SetChatPhotoParams https://core.telegram.org/bots/api#setchatphoto
type SetChatPhotoParams struct {
	ChatID any              `json:"chat_id"`
	Photo  *types.InputFile `json:"photo,omitempty" media:"photo"` // FIXME Id и URL - не работает загрузка
}

// DeleteChatPhotoParams https://core.telegram.org/bots/api#deletechatphoto
type DeleteChatPhotoParams struct {
	ChatID any `json:"chat_id"`
}

// SetChatTitleParams https://core.telegram.org/bots/api#setchattitle
type SetChatTitleParams struct {
	ChatID any    `json:"chat_id"`
	Title  string `json:"title"`
}

// SetChatDescriptionParams https://core.telegram.org/bots/api#setchatdescription
type SetChatDescriptionParams struct {
	ChatID      any    `json:"chat_id"`
	Description string `json:"description"`
}

// PinChatMessageParams https://core.telegram.org/bots/api#pinchatmessage
type PinChatMessageParams struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageID            int    `json:"message_id"`
	DisableNotification  bool   `json:"disable_notification,omitempty"`
}

// UnpinChatMessageParams https://core.telegram.org/bots/api#unpinchatmessage
type UnpinChatMessageParams struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	ChatID               any    `json:"chat_id"`
	MessageID            int    `json:"message_id,omitempty"`
}

// UnpinAllChatMessagesParams https://core.telegram.org/bots/api#unpinallchatmessages
type UnpinAllChatMessagesParams struct {
	ChatID any `json:"chat_id"`
}

// LeaveChatParams https://core.telegram.org/bots/api#leavechat
type LeaveChatParams struct {
	ChatID any `json:"chat_id"`
}

// GetChatParams https://core.telegram.org/bots/api#getchat
type GetChatParams struct {
	ChatID any `json:"chat_id"`
}

// GetChatAdministratorsParams https://core.telegram.org/bots/api#getchatadministrators
type GetChatAdministratorsParams struct {
	ChatID any `json:"chat_id"`
}

// GetChatMemberCountParams https://core.telegram.org/bots/api#getchatmembercount
type GetChatMemberCountParams struct {
	ChatID any `json:"chat_id"`
}

// GetChatMemberParams https://core.telegram.org/bots/api#getchatmember
type GetChatMemberParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

// SetChatStickerSetParams https://core.telegram.org/bots/api#setchatstickerset
type SetChatStickerSetParams struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

type DeleteChatStickerSetParams struct {
	ChatID any `json:"chat_id"`
}

// CreateForumTopicParams https://core.telegram.org/bots/api#createforumtopic
type CreateForumTopicParams struct {
	ChatID            any    `json:"chat_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// EditForumTopicParams https://core.telegram.org/bots/api#editforumtopic
type EditForumTopicParams struct {
	ChatID            any    `json:"chat_id"`
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// CloseForumTopicParams https://core.telegram.org/bots/api#closeforumtopic
type CloseForumTopicParams struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

// ReopenForumTopicParams https://core.telegram.org/bots/api#reopenforumtopic
type ReopenForumTopicParams struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

// DeleteForumTopicParams https://core.telegram.org/bots/api#deleteforumtopic
type DeleteForumTopicParams struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

// UnpinAllForumTopicMessagesParams https://core.telegram.org/bots/api#unpinallforumtopicmessages
type UnpinAllForumTopicMessagesParams struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

// EditGeneralForumTopicParams https://core.telegram.org/bots/api#editgeneralforumtopic
type EditGeneralForumTopicParams struct {
	ChatID any    `json:"chat_id"`
	Name   string `json:"name"`
}

// CloseGeneralForumTopicParams https://core.telegram.org/bots/api#closegeneralforumtopic
type CloseGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

// ReopenGeneralForumTopicParams https://core.telegram.org/bots/api#reopengeneralforumtopic
type ReopenGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

// HideGeneralForumTopicParams https://core.telegram.org/bots/api#hidegeneralforumtopic
type HideGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

// UnhideGeneralForumTopicParams https://core.telegram.org/bots/api#unhidegeneralforumtopic
type UnhideGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

// UnpinAllGeneralForumTopicMessagesParams https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
type UnpinAllGeneralForumTopicMessagesParams struct {
	ChatID any `json:"chat_id"`
}

// AnswerCallbackQueryParams https://core.telegram.org/bots/api#answercallbackquery
type AnswerCallbackQueryParams struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

// GetUserChatBoostsParams https://core.telegram.org/bots/api#getuserchatboosts
type GetUserChatBoostsParams struct {
	ChatID any `json:"chat_id"`
	UserID int `json:"user_id"`
}

// GetBusinessConnectionParams https://core.telegram.org/bots/api#getbusinessconnection
type GetBusinessConnectionParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

// SetMyCommandsParams https://core.telegram.org/bots/api#setmycommands
type SetMyCommandsParams struct {
	Commands     []types.BotCommand    `json:"commands"`
	Scope        types.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                `json:"language_code,omitempty"`
}

// DeleteMyCommandsParams https://core.telegram.org/bots/api#deletemycommands
type DeleteMyCommandsParams struct {
	Scope        types.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                `json:"language_code,omitempty"`
}

// GetMyCommandsParams https://core.telegram.org/bots/api#getmycommands
type GetMyCommandsParams struct {
	Scope        types.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                `json:"language_code,omitempty"`
}

// SetMyNameParams https://core.telegram.org/bots/api#setmyname
type SetMyNameParams struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyNameParams https://core.telegram.org/bots/api#getmyname
type GetMyNameParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyDescriptionParams https://core.telegram.org/bots/api#setmydescription
type SetMyDescriptionParams struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyDescriptionParams https://core.telegram.org/bots/api#getmydescription
type GetMyDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyShortDescriptionParams https://core.telegram.org/bots/api#setmyshortdescription
type SetMyShortDescriptionParams struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

// GetMyShortDescriptionParams https://core.telegram.org/bots/api#getmyshortdescription
type GetMyShortDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// SetChatMenuButtonParams https://core.telegram.org/bots/api#setchatmenubutton
type SetChatMenuButtonParams struct {
	ChatID     any              `json:"chat_id,omitempty"`
	MenuButton types.MenuButton `json:"menu_button,omitempty"`
}

// GetChatMenuButtonParams https://core.telegram.org/bots/api#getchatmenubutton
type GetChatMenuButtonParams struct {
	ChatID any `json:"chat_id,omitempty"`
}

// SetMyDefaultAdministratorRightsParams https://core.telegram.org/bots/api#setmydefaultadministratorrights
type SetMyDefaultAdministratorRightsParams struct {
	Rights      *types.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                           `json:"for_channels,omitempty"`
}

// GetMyDefaultAdministratorRightsParams https://core.telegram.org/bots/api#getmydefaultadministratorrights
type GetMyDefaultAdministratorRightsParams struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

// SendGiftParams https://core.telegram.org/bots/api#sendgift
type SendGiftParams struct {
	UserID        int64                 `json:"user_id,omitempty"`
	ChatID        any                   `json:"chat_id,omitempty"`
	GiftID        string                `json:"gift_id"`
	PayForUpgrade bool                  `json:"pay_for_upgrade,omitempty"`
	Text          string                `json:"text,omitempty"`
	TextParseMode types.ParseMode       `json:"text_parse_mode,omitempty"`
	TextEntities  []types.MessageEntity `json:"text_entities,omitempty"`
}

// GiftPremiumSubscriptionParams https://core.telegram.org/bots/api#giftpremiumsubscription
type GiftPremiumSubscriptionParams struct {
	UserID        int64                 `json:"user_id"`
	MonthCount    int                   `json:"month_count"`
	StarCount     int                   `json:"star_count"`
	Text          string                `json:"text,omitempty"`
	TextParseMode types.ParseMode       `json:"text_parse_mode,omitempty"`
	TextEntities  []types.MessageEntity `json:"text_entities,omitempty"`
}

// VerifyUserParams https://core.telegram.org/bots/api#verifyuser
type VerifyUserParams struct {
	UserID            int64  `json:"user_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

// VerifyChatParams https://core.telegram.org/bots/api#verifychat
type VerifyChatParams struct {
	ChatID            any    `json:"chat_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

// RemoveUserVerificationParams https://core.telegram.org/bots/api#removeuserverification
type RemoveUserVerificationParams struct {
	UserID int64 `json:"user_id"`
}

// RemoveChatVerificationParams https://core.telegram.org/bots/api#removechatverification
type RemoveChatVerificationParams struct {
	ChatID any `json:"chat_id"`
}

// ReadBusinessMessageParams https://core.telegram.org/bots/api#readbusinessmessage
type ReadBusinessMessageParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	ChatID               int    `json:"chat_id"`
	MessageID            int    `json:"message_id"`
}

// DeleteBusinessMessagesParams https://core.telegram.org/bots/api#deletebusinessmessages
type DeleteBusinessMessagesParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	MessageIDs           []int  `json:"message_ids"`
}

// SetBusinessAccountNameParams https://core.telegram.org/bots/api#setbusinessaccountname
type SetBusinessAccountNameParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name,omitempty"`
}

// SetBusinessAccountUsernameParams https://core.telegram.org/bots/api#setbusinessaccountusername
type SetBusinessAccountUsernameParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Username             string `json:"username,omitempty"`
}

// SetBusinessAccountBioParams https://core.telegram.org/bots/api#setbusinessaccountbio
type SetBusinessAccountBioParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Bio                  string `json:"bio,omitempty"`
}

// SetBusinessAccountProfilePhotoParams https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
type SetBusinessAccountProfilePhotoParams struct {
	BusinessConnectionID string                  `json:"business_connection_id"`
	Photo                types.InputProfilePhoto `json:"photo"`
	IsPublic             bool                    `json:"is_public,omitempty"`
}

// RemoveBusinessAccountProfilePhotoParams https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
type RemoveBusinessAccountProfilePhotoParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	IsPublic             bool   `json:"is_public,omitempty"`
}

// SetBusinessAccountGiftSettingsParams https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
type SetBusinessAccountGiftSettingsParams struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	ShowGiftButton       bool                     `json:"show_gift_button"`
	AcceptedGiftTypes    *types.AcceptedGiftTypes `json:"accepted_gift_types"`
}

// GetBusinessAccountStarBalanceParams https://core.telegram.org/bots/api#getbusinessaccountstarbalance
type GetBusinessAccountStarBalanceParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

// TransferBusinessAccountStarsParams https://core.telegram.org/bots/api#transferbusinessaccountstars
type TransferBusinessAccountStarsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StarCount            int    `json:"star_count"`
}

// GetBusinessAccountGiftsParams https://core.telegram.org/bots/api#getbusinessaccountgifts
type GetBusinessAccountGiftsParams struct {
	BusinessConnectionID        string `json:"business_connection_id"`
	ExcludeUnsaved              bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int    `json:"limit,omitempty"`
}

// GetUserGiftsParams https://core.telegram.org/bots/api#getusergifts
type GetUserGiftsParams struct {
	UserID                      int64  `json:"user_id"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int    `json:"limit,omitempty"`
}

// GetChatGiftsParams https://core.telegram.org/bots/api#getchatgifts
type GetChatGiftsParams struct {
	ChatID                      any    `json:"chat_id"`
	ExcludeUnsaved              bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeFromBlockchain       bool   `json:"exclude_from_blockchain,omitempty"`
	ExcludeUnique               bool   `json:"exclude_unique,omitempty"`
	SortByPrice                 bool   `json:"sort_by_price,omitempty"`
	Offset                      string `json:"offset,omitempty"`
	Limit                       int    `json:"limit,omitempty"`
}

// ConvertGiftToStarsParams https://core.telegram.org/bots/api#convertgifttostars
type ConvertGiftToStarsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
}

// UpgradeGiftParams https://core.telegram.org/bots/api#upgradegift
type UpgradeGiftParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	KeepOriginalDetails  bool   `json:"keep_original_details,omitempty"`
	StarCount            int    `json:"star_count,omitempty"`
}

// TransferGiftParams https://core.telegram.org/bots/api#transfergift
type TransferGiftParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	NewOwnerChatID       int    `json:"new_owner_chat_id"`
	StarCount            int    `json:"star_count"`
}

// PostStoryParams https://core.telegram.org/bots/api#poststory
type PostStoryParams struct {
	BusinessConnectionID string                  `json:"business_connection_id"`
	Content              types.InputStoryContent `json:"content"`
	ActivePeriod         int                     `json:"active_period"`
	Caption              string                  `json:"caption,omitempty"`
	ParseMode            types.ParseMode         `json:"parse_mode,omitempty"`
	CaptionEntities      []types.MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []types.StoryArea       `json:"areas,omitempty"`
	PostToChatPage       bool                    `json:"post_to_chat_page,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
}

// RepostStoryParams https://core.telegram.org/bots/api#repoststory
type RepostStoryParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	FromChatID           any    `json:"from_chat_id"`
	FromStoryID          int    `json:"from_story_id"`
	ActivePeriod         int    `json:"active_period"`
	PostToChatPage       bool   `json:"post_to_chat_page,omitempty"`
	ProtectContent       bool   `json:"protect_content,omitempty"`
}

// EditStoryParams https://core.telegram.org/bots/api#editstory
type EditStoryParams struct {
	BusinessConnectionID string                  `json:"business_connection_id"`
	StoryID              int                     `json:"story_id"`
	Content              types.InputStoryContent `json:"content"`
	Caption              string                  `json:"caption,omitempty"`
	ParseMode            types.ParseMode         `json:"parse_mode,omitempty"`
	CaptionEntities      []types.MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []types.StoryArea       `json:"areas,omitempty"`
}

// DeleteStoryParams https://core.telegram.org/bots/api#deletestory
type DeleteStoryParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StoryID              int    `json:"story_id"`
}

// EditMessageTextParams https://core.telegram.org/bots/api#editmessagetext
type EditMessageTextParams struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               any                         `json:"chat_id,omitempty"`
	MessageID            int                         `json:"message_id,omitempty"`
	InlineMessageID      string                      `json:"inline_message_id,omitempty"`
	Text                 string                      `json:"text"`
	ParseMode            types.ParseMode             `json:"parse_mode,omitempty"`
	Entities             []types.MessageEntity       `json:"entities,omitempty"`
	LinkPreviewOptions   *types.LinkPreviewOptions   `json:"link_preview_options,omitempty"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageCaptionParams https://core.telegram.org/bots/api#editmessagecaption
type EditMessageCaptionParams struct {
	BusinessConnectionID  string                      `json:"business_connection_id,omitempty"`
	ChatID                any                         `json:"chat_id,omitempty"`
	MessageID             int                         `json:"message_id,omitempty"`
	InlineMessageID       string                      `json:"inline_message_id,omitempty"`
	Caption               string                      `json:"caption,omitempty"`
	ParseMode             types.ParseMode             `json:"parse_mode,omitempty"`
	CaptionEntities       []types.MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                        `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageMediaParams https://core.telegram.org/bots/api#editmessagemedia
type EditMessageMediaParams struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               any                         `json:"chat_id,omitempty"`
	MessageID            int                         `json:"message_id,omitempty"`
	InlineMessageID      string                      `json:"inline_message_id,omitempty"`
	Media                types.InputMedia            `json:"media"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocationParams https://core.telegram.org/bots/api#editmessagelivelocation
type EditMessageLiveLocationParams struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               any                         `json:"chat_id,omitempty"`
	MessageID            int                         `json:"message_id,omitempty"`
	InlineMessageID      string                      `json:"inline_message_id,omitempty"`
	Latitude             float64                     `json:"latitude"`
	Longitude            float64                     `json:"longitude"`
	LivePeriod           int                         `json:"live_period,omitempty"`
	HorizontalAccuracy   float64                     `json:"horizontal_accuracy,omitempty"`
	Heading              int                         `json:"heading,omitempty"`
	ProximityAlertRadius int                         `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocationParams https://core.telegram.org/bots/api#stopmessagelivelocation
type StopMessageLiveLocationParams struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               any                         `json:"chat_id,omitempty"`
	MessageID            int                         `json:"message_id,omitempty"`
	InlineMessageID      string                      `json:"inline_message_id,omitempty"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageChecklistParams https://core.telegram.org/bots/api#editmessagechecklist
type EditMessageChecklistParams struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               int                         `json:"chat_id,omitempty"`
	MessageID            int                         `json:"message_id,omitempty"`
	Checklist            *types.InputChecklist       `json:"checklist"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkupParams https://core.telegram.org/bots/api#editmessagereplymarkup
type EditMessageReplyMarkupParams struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               any                         `json:"chat_id,omitempty"`
	MessageID            int                         `json:"message_id,omitempty"`
	InlineMessageID      string                      `json:"inline_message_id,omitempty"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopPollParams https://core.telegram.org/bots/api#stoppoll
type StopPollParams struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               any                         `json:"chat_id"`
	MessageID            int                         `json:"message_id"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// ApproveSuggestedPostParams https://core.telegram.org/bots/api#approvesuggestedpost
type ApproveSuggestedPostParams struct {
	ChatID    int `json:"chat_id"`
	MessageID int `json:"message_id"`
	SendDate  int `json:"send_date,omitempty"`
}

// DeclineSuggestedPostParams https://core.telegram.org/bots/api#declinesuggestedpost
type DeclineSuggestedPostParams struct {
	ChatID    int    `json:"chat_id"`
	MessageID int    `json:"message_id"`
	Comment   string `json:"comment,omitempty"`
}

// DeleteMessageParams https://core.telegram.org/bots/api#deletemessage
type DeleteMessageParams struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id"`
}

// DeleteMessagesParams https://core.telegram.org/bots/api#deletemessages
type DeleteMessagesParams struct {
	ChatID     any   `json:"chat_id"`
	MessageIDs []int `json:"message_ids"`
}

// SendInvoiceParams https://core.telegram.org/bots/api#sendinvoice
type SendInvoiceParams struct {
	ChatID                    any                            `json:"chat_id"`
	MessageThreadID           int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID     int                            `json:"direct_messages_topic_id,omitempty"`
	Title                     string                         `json:"title"`
	Description               string                         `json:"description"`
	Payload                   string                         `json:"payload"`
	ProviderToken             string                         `json:"provider_token,omitempty"`
	Currency                  string                         `json:"currency"`
	Prices                    []types.LabeledPrice           `json:"prices"`
	MaxTipAmount              int                            `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                          `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                         `json:"start_parameter,omitempty"`
	ProviderData              string                         `json:"provider_data,omitempty"`
	PhotoURL                  string                         `json:"photo_url,omitempty"`
	PhotoSize                 int                            `json:"photo_size,omitempty"`
	PhotoWidth                int                            `json:"photo_width,omitempty"`
	PhotoHeight               int                            `json:"photo_height,omitempty"`
	NeedName                  bool                           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                           `json:"need_email,omitempty"`
	NeedShippingAddress       bool                           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                           `json:"is_flexible,omitempty"`
	DisableNotification       bool                           `json:"disable_notification,omitempty"`
	ProtectContent            bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast        bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID           string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters   *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters           *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup               *types.InlineKeyboardMarkup    `json:"reply_markup,omitempty"`
}

// CreateInvoiceLinkParams https://core.telegram.org/bots/api#createinvoicelink
type CreateInvoiceLinkParams struct {
	BusinessConnectionID      string               `json:"business_connection_id,omitempty"`
	Title                     string               `json:"title"`
	Description               string               `json:"description"`
	Payload                   string               `json:"payload"`
	ProviderToken             string               `json:"provider_token,omitempty"`
	Currency                  string               `json:"currency"`
	Prices                    []types.LabeledPrice `json:"prices"`
	SubscriptionPeriod        int                  `json:"subscription_period,omitempty"`
	MaxTipAmount              int                  `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string               `json:"provider_data,omitempty"`
	PhotoURL                  string               `json:"photo_url,omitempty"`
	PhotoSize                 int                  `json:"photo_size,omitempty"`
	PhotoWidth                int                  `json:"photo_width,omitempty"`
	PhotoHeight               int                  `json:"photo_height,omitempty"`
	NeedName                  bool                 `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                 `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                 `json:"need_email,omitempty"`
	NeedShippingAddress       bool                 `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                 `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                 `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                 `json:"is_flexible,omitempty"`
}

// AnswerShippingQueryParams https://core.telegram.org/bots/api#answershippingquery
type AnswerShippingQueryParams struct {
	ShippingQueryID string                 `json:"shipping_query_id"`
	OK              bool                   `json:"ok"`
	ShippingOptions []types.ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string                 `json:"error_message,omitempty"`
}

// AnswerPreCheckoutQueryParams https://core.telegram.org/bots/api#answerprecheckoutquery
type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	OK                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

// GetStarTransactionsParams https://core.telegram.org/bots/api#getstartransactions
type GetStarTransactionsParams struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// RefundStarPaymentParams https://core.telegram.org/bots/api#refundstarpayment
type RefundStarPaymentParams struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
}

// EditUserStarSubscriptionParams https://core.telegram.org/bots/api#edituserstarsubscription
type EditUserStarSubscriptionParams struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	IsCanceled              bool   `json:"is_canceled"`
}

type SetPassportDataErrorsParams struct {
	UserID int64                        `json:"user_id"`
	Errors []types.PassportElementError `json:"errors"`
}

// SendStickerParams https://core.telegram.org/bots/api#sendsticker
type SendStickerParams struct {
	BusinessConnectionID    string                         `json:"business_connection_id,omitempty"`
	ChatID                  any                            `json:"chat_id"`
	MessageThreadID         int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                            `json:"direct_messages_topic_id,omitempty"`
	Sticker                 *types.InputFile               `json:"sticker" media:"sticker"`
	Emoji                   string                         `json:"emoji,omitempty"`
	DisableNotification     bool                           `json:"disable_notification,omitempty"`
	ProtectContent          bool                           `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                           `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                         `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *types.SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *types.ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             types.ReplyMarkup              `json:"reply_markup,omitempty"`
}

// GetStickerSetParams https://core.telegram.org/bots/api#getstickerset
type GetStickerSetParams struct {
	Name string `json:"name"`
}

// GetCustomEmojiStickersParams https://core.telegram.org/bots/api#getcustomemojistickers
type GetCustomEmojiStickersParams struct {
	CustomEmojiIDs []string `json:"custom_emoji_ids"`
}

// UploadStickerFileParams https://core.telegram.org/bots/api#uploadstickerfile
type UploadStickerFileParams struct {
	UserID        int64            `json:"user_id"`
	Sticker       *types.InputFile `json:"sticker" media:"sticker"`
	StickerFormat string           `json:"sticker_format"`
}

// CreateNewStickerSetParams https://core.telegram.org/bots/api#createnewstickerset
type CreateNewStickerSetParams struct {
	UserID          int64                `json:"user_id"`
	Name            string               `json:"name"`
	Title           string               `json:"title"`
	Stickers        []types.InputSticker `json:"stickers"`
	StickerType     string               `json:"sticker_type,omitempty"`
	NeedsRepainting bool                 `json:"needs_repainting,omitempty"`
}

// AddStickerToSetParams https://core.telegram.org/bots/api#addstickertoset
type AddStickerToSetParams struct {
	UserID  int64               `json:"user_id"`
	Name    string              `json:"name"`
	Sticker *types.InputSticker `json:"sticker"`
}

// SetStickerPositionInSetParams https://core.telegram.org/bots/api#setstickerpositioninset
type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

// DeleteStickerFromSetParams https://core.telegram.org/bots/api#deletestickerfromset
type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"`
}

// ReplaceStickerInSetParams https://core.telegram.org/bots/api#replacestickerinset
type ReplaceStickerInSetParams struct {
	UserID     int64               `json:"user_id"`
	Name       string              `json:"name"`
	OldSticker string              `json:"old_sticker"`
	Sticker    *types.InputSticker `json:"sticker"`
}

// SetStickerEmojiListParams https://core.telegram.org/bots/api#setstickeremojilist
type SetStickerEmojiListParams struct {
	Sticker   string   `json:"sticker"`
	EmojiList []string `json:"emoji_list"`
}

// SetStickerKeywordsParams https://core.telegram.org/bots/api#setstickerkeywords
type SetStickerKeywordsParams struct {
	Sticker  string   `json:"sticker"`
	Keywords []string `json:"keywords,omitempty"`
}

// SetStickerMaskPositionParams https://core.telegram.org/bots/api#setstickermaskposition
type SetStickerMaskPositionParams struct {
	Sticker      string              `json:"sticker"`
	MaskPosition *types.MaskPosition `json:"mask_position,omitempty"`
}

// SetStickerSetTitleParams https://core.telegram.org/bots/api#setstickersettitle
type SetStickerSetTitleParams struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

// SetStickerSetThumbnailParams https://core.telegram.org/bots/api#setstickersetthumbnail
type SetStickerSetThumbnailParams struct {
	Name      string           `json:"name"`
	UserID    int64            `json:"user_id"`
	Thumbnail *types.InputFile `json:"thumbnail,omitempty" media:"thumbnail"`
	Format    string           `json:"format"`
}

// SetCustomEmojiStickerSetThumbnailParams https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
type SetCustomEmojiStickerSetThumbnailParams struct {
	Name          string `json:"name"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// DeleteStickerSetParams https://core.telegram.org/bots/api#deletestickerset
type DeleteStickerSetParams struct {
	Name string `json:"name"`
}

// AnswerInlineQueryParams https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQueryParams struct {
	InlineQueryID string                          `json:"inline_query_id"`
	Results       []types.InlineQueryResult       `json:"results"`
	CacheTime     int                             `json:"cache_time,omitempty"`
	IsPersonal    bool                            `json:"is_personal,omitempty"`
	NextOffset    string                          `json:"next_offset,omitempty"`
	Button        *types.InlineQueryResultsButton `json:"button,omitempty"`
}

// AnswerWebAppQueryParams https://core.telegram.org/bots/api#answerwebappquery
type AnswerWebAppQueryParams struct {
	WebAppQueryID string                  `json:"web_app_query_id"`
	Result        types.InlineQueryResult `json:"result"`
}

// SavePreparedInlineMessageParams https://core.telegram.org/bots/api#savepreparedinlinemessage
type SavePreparedInlineMessageParams struct {
	UserID            int64                   `json:"user_id"`
	Result            types.InlineQueryResult `json:"result"`
	AllowUserChats    bool                    `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool                    `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool                    `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool                    `json:"allow_channel_chats,omitempty"`
}

// SendGameParams https://core.telegram.org/bots/api#sendgame
type SendGameParams struct {
	BusinessConnectionID string                 `json:"business_connection_id,omitempty"`
	ChatID               any                    `json:"chat_id"`
	MessageThreadID      int                    `json:"message_thread_id,omitempty"`
	GameShorName         string                 `json:"game_short_name"`
	DisableNotification  bool                   `json:"disable_notification,omitempty"`
	ProtectContent       bool                   `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                   `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                 `json:"message_effect_id,omitempty"`
	ReplyParameters      *types.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          types.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SetGameScoreParams https://core.telegram.org/bots/api#setgamescore
type SetGameScoreParams struct {
	UserID             int64 `json:"user_id"`
	Score              int   `json:"score"`
	Force              bool  `json:"force,omitempty"`
	DisableEditMessage bool  `json:"disable_edit_message,omitempty"`
	ChatID             any   `json:"chat_id,omitempty"`
	MessageID          int   `json:"message_id,omitempty"`
	InlineMessageID    int   `json:"inline_message_id,omitempty"`
}

type GetGameHighScoresParams struct {
	UserID          int64 `json:"user_id"`
	ChatID          any   `json:"chat_id,omitempty"`
	MessageID       int   `json:"message_id,omitempty"`
	InlineMessageID int   `json:"inline_message_id,omitempty"`
}

// SetWebhookParams https://core.telegram.org/bots/api#setwebhook
type SetWebhookParams struct {
	URL                string               `json:"url"`
	Certificate        *types.InputFile     `json:"certificate,omitempty" media:"certificate"`
	IPAddress          string               `json:"ip_address,omitempty"`
	MaxConnections     int                  `json:"max_connections,omitempty"`
	AllowedUpdates     types.AllowedUpdates `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool                 `json:"drop_pending_updates,omitempty"`
	SecretToken        string               `json:"secret_token,omitempty"`
}

// DeleteWebhookParams https://core.telegram.org/bots/api#deletewebhook
type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}
