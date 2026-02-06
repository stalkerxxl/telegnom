package types

// InlineQueryResult represents one result of an inline query. Telegram clients
// currently support results of the following 20 types.
//
//	Note: All URLs passed in nline query results will be available to end users
//	and therefore must be assumed to be public.
//
// See https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult interface {
	isInlineQueryResult()
}

// Reserved for future input types
// InlineQueryResultData struct {
//	impl InlineQueryResult
//}

// InlineQueryResultArticle represents a link to an article or web page.
//
// See https://core.telegram.org/bots/api#inlinequeryresultarticle
type InlineQueryResultArticle struct {
	Type                string              `json:"type"` // must be "article"
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	InputMessageContent InputMessageContent `json:"input_message_content"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	URL                 string              `json:"url,omitempty"`
	Description         string              `json:"description,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

func (iqr *InlineQueryResultArticle) isInlineQueryResult() {}

// InlineQueryResultPhoto represents a link to a photo. By default, this photo
// will be sent by the user with optional caption. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the photo.
//
// See https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	Type                  string              `json:"type"` // must be "photo"
	ID                    string              `json:"id"`
	PhotoURL              string              `json:"photo_url"`
	ThumbnailURL          string              `json:"thumbnail_url"`
	PhotoWidth            int                 `json:"photo_width,omitempty"`
	PhotoHeight           int                 `json:"photo_height,omitempty"`
	Title                 string              `json:"title,omitempty"`
	Description           string              `json:"description,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultPhoto) isInlineQueryResult() {}

// InlineQueryResultGif represents a link to an animated GIF file. By default,
// this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the
// specified content instead of the animation.
//
// See https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGif struct {
	Type                  string              `json:"type"` // must be "gif"
	ID                    string              `json:"id"`
	GifURL                string              `json:"gif_url"`
	GifWidth              int                 `json:"gif_width,omitempty"`
	GifHeight             int                 `json:"gif_height,omitempty"`
	GifDuration           int                 `json:"gif_duration,omitempty"`
	ThumbnailURL          string              `json:"thumbnail_url"`
	ThumbnailMimeType     string              `json:"thumbnail_mime_type,omitempty"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultGif) isInlineQueryResult() {}

// InlineQueryResultMpeg4Gif represents a link to a video animation (H.264/MPEG-4
// AVC video without sound). By default, this animated MPEG-4 file will be sent
// by the user with optional caption. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the animation.
//
// See https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMpeg4Gif struct {
	Type                  string              `json:"type"` // must be "mpeg4_gif"
	ID                    string              `json:"id"`
	Mpeg4URL              string              `json:"mpeg4_url"`
	Mpeg4Width            int                 `json:"mpeg4_width,omitempty"`
	Mpeg4Height           int                 `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         int                 `json:"mpeg4_duration,omitempty"`
	ThumbnailURL          string              `json:"thumbnail_url"`
	ThumbnailMimeType     string              `json:"thumbnail_mime_type,omitempty"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultMpeg4Gif) isInlineQueryResult() {}

// InlineQueryResultVideo Represents a link to a page containing an embedded
// video player or a video file. By default, this video file will be sent by the
// user with an optional caption. Alternatively, you can use InputMessageContent
// to send a message with the specified content instead of the video.
//
//	If an InlineQueryResultVideo message contains an embedded video (e.g.,
//	YouTube), you must replace its content using InputMessageContent.
//
// See https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	Type                  string              `json:"type"` // must be "video"
	ID                    string              `json:"id"`
	VideoURL              string              `json:"video_url"`
	MimeType              string              `json:"mime_type"`
	ThumbnailURL          string              `json:"thumbnail_url"`
	Title                 string              `json:"title"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	VideoWidth            int                 `json:"video_width,omitempty"`
	VideoHeight           int                 `json:"video_height,omitempty"`
	VideoDuration         int                 `json:"video_duration,omitempty"`
	Description           string              `json:"description,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultVideo) isInlineQueryResult() {}

// InlineQueryResultAudio represents a link to an MP3 audio file. By default,
// this audio file will be sent by the user. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the audio.
//
// See https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	Type                string              `json:"type"` // must be "audio"
	ID                  string              `json:"id"`
	AudioURL            string              `json:"audio_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	Performer           string              `json:"performer,omitempty"`
	AudioDuration       int                 `json:"audio_duration,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultAudio) isInlineQueryResult() {}

// InlineQueryResultVoice represents a link to a voice recording in an .OGG
// container encoded with OPUS. By default, this voice recording will be sent by
// the user. Alternatively, you can use InputMessageContent to send a message
// with the specified content instead of the voice message.
//
// See https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	Type                string              `json:"type"` // must be "voice"
	ID                  string              `json:"id"`
	VoiceURL            string              `json:"voice_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	VoiceDuration       int                 `json:"voice_duration,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultVoice) isInlineQueryResult() {}

// InlineQueryResultDocument represents a link to a file. By default, this file
// will be sent by the user with an optional caption. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the file. Currently, only .PDF and .ZIP files can be sent using this method.
//
// See https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	Type                string              `json:"type"` // must be "document"
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	DocumentURL         string              `json:"document_url"`
	MimeType            string              `json:"mime_type"`
	Description         string              `json:"description,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

func (iqr *InlineQueryResultDocument) isInlineQueryResult() {}

// InlineQueryResultLocation represents a location on a map. By default, the
// location will be sent by the user. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the location.
//
// See https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	Type                 string              `json:"type"` // must be "location"
	ID                   string              `json:"id"`
	Latitude             float64             `json:"latitude"`
	Longitude            float64             `json:"longitude"`
	Title                string              `json:"title"`
	HorizontalAccuracy   float64             `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                 `json:"live_period,omitempty"`
	Heading              int                 `json:"heading,omitempty"`
	ProximityAlertRadius int                 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent  InputMessageContent `json:"input_message_content,omitempty"`
	ThumbnailURL         string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      int                 `json:"thumbnail_height,omitempty"`
}

func (iqr *InlineQueryResultLocation) isInlineQueryResult() {}

// InlineQueryResultVenue Represents a venue. By default, the venue will be sent
// by the user. Alternatively, you can use InputMessageContent to send a
// message with the specified content instead of the venue.
//
// See https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	Type                string              `json:"type"` // must be "venue"
	ID                  string              `json:"id"`
	Latitude            float64             `json:"latitude"`
	Longitude           float64             `json:"longitude"`
	Title               string              `json:"title"`
	Address             string              `json:"address"`
	FoursquareID        string              `json:"foursquare_id,omitempty"`
	FoursquareType      string              `json:"foursquare_type,omitempty"`
	GooglePlaceID       string              `json:"google_place_id,omitempty"`
	GooglePlaceType     string              `json:"google_place_type,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

func (iqr *InlineQueryResultVenue) isInlineQueryResult() {}

// InlineQueryResultContact represents a contact with a phone number. By default,
// this contact will be sent by the user. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the contact.
//
// See https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	Type                string              `json:"type"` // must be "contact"
	ID                  string              `json:"id"`
	PhoneNumber         string              `json:"phone_number"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name,omitempty"`
	VCard               string              `json:"vcard,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

func (iqr *InlineQueryResultContact) isInlineQueryResult() {}

// InlineQueryResultGame represents a Game.
//
// See https://core.telegram.org/bots/api#inlinequeryresultgame
type InlineQueryResultGame struct {
	Type          string      `json:"type"` // must be "game"
	ID            string      `json:"id"`
	GameShortName string      `json:"game_short_name"`
	ReplyMarkup   ReplyMarkup `json:"reply_markup,omitempty"`
}

func (iqr *InlineQueryResultGame) isInlineQueryResult() {}

// InlineQueryResultCachedPhoto represents a link to a photo stored on the
// Telegram servers. By default, this photo will be sent by the user with an
// optional caption. Alternatively, you can use InputMessageContent to send a
// message with the specified content instead of the photo.
//
// See https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	Type                  string              `json:"type"` // must be "photo"
	ID                    string              `json:"id"`
	PhotoFileID           string              `json:"photo_file_id"`
	Title                 string              `json:"title,omitempty"`
	Description           string              `json:"description,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultCachedPhoto) isInlineQueryResult() {}

// InlineQueryResultCachedGif represents a link to an animated GIF file stored on
// the Telegram servers. By default, this animated GIF file will be sent by the
// user with an optional caption. Alternatively, you can use
// InputMessageContent to send a message with specified content instead of the
// animation.
//
// See https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGif struct {
	Type                  string              `json:"type"` // must be "gif"
	ID                    string              `json:"id"`
	GifFileID             string              `json:"gif_file_id"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultCachedGif) isInlineQueryResult() {}

// InlineQueryResultCachedMpeg4Gif represents a link to a video animation
// (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By
// default, this animated MPEG-4 file will be sent by the user with an optional
// caption. Alternatively, you can use InputMessageContent to send a message
// with the specified content instead of the animation.
//
// See https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMpeg4Gif struct {
	Type                  string              `json:"type"` // must be "mpeg4_gif"
	ID                    string              `json:"id"`
	Mpeg4FileID           string              `json:"mpeg4_file_id"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultCachedMpeg4Gif) isInlineQueryResult() {}

// InlineQueryResultCachedSticker represents a link to a sticker stored on the
// Telegram servers. By default, this sticker will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the
// specified content instead of the sticker.
//
// See https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	Type                string              `json:"type"` // must be "sticker"
	ID                  string              `json:"id"`
	StickerFileID       string              `json:"sticker_file_id"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultCachedSticker) isInlineQueryResult() {}

// InlineQueryResultCachedDocument represents a link to a file stored on the
// Telegram servers. By default, this file will be sent by the user with an
// optional caption. Alternatively, you can use InputMessageContent to send a
// message with the specified content instead of the file.
//
// See https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	Type                string              `json:"type"` // must be "document"
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	DocumentFileID      string              `json:"document_file_id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultCachedDocument) isInlineQueryResult() {}

// InlineQueryResultCachedVideo represents a link to a video file stored on the
// Telegram servers. By default, this video file will be sent by the user with an
// optional caption. Alternatively, you can use InputMessageContent to send a
// message with the specified content instead of the video.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	Type                  string              `json:"type"` // must be "video"
	ID                    string              `json:"id"`
	VideoFileID           string              `json:"video_file_id"`
	Title                 string              `json:"title"`
	Description           string              `json:"description,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultCachedVideo) isInlineQueryResult() {}

// InlineQueryResultCachedVoice represents a link to a voice message stored on
// the Telegram servers. By default, this voice message will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the
// specified content instead of the voice message.
//
// See https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	Type                string              `json:"type"` // must be "voice"
	ID                  string              `json:"id"`
	VoiceFileID         string              `json:"voice_file_id"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultCachedVoice) isInlineQueryResult() {}

// InlineQueryResultCachedAudio represents a link to an MP3 audio file stored on
// the Telegram servers. By default, this audio file will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the
// specified content instead of the audio.
//
// See https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	Type                string              `json:"type"` // must be "audio"
	ID                  string              `json:"id"`
	AudioFileID         string              `json:"audio_file_id"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (iqr *InlineQueryResultCachedAudio) isInlineQueryResult() {}
