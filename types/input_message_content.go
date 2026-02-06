package types

// InputMessageContent represents the content of a message to be sent as a result
// of an inline query. Currently, only the following 5 types are supported:
// InputTextMessageContent, InputLocationMessageContent,
// InputVenueMessageContent, InputContactMessageContent, and
// InputInvoiceMessageContent.
//
// See https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent interface {
	isInputMessageContent()
}

//// InputMessageContentData reserved for future use
//type InputMessageContentData struct {
//	impl InputMessageContent
//}

// InputTextMessageContent represents the content of a text message to be sent as
// the result of an inline query.
//
// See https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          ParseMode           `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

func (imc *InputTextMessageContent) isInputMessageContent() {}

// InputLocationMessageContent represents the content of a location message to be
// sent as the result of an inline query.
//
// See https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

func (imc *InputLocationMessageContent) isInputMessageContent() {}

// InputVenueMessageContent represents the content of a venue message to be sent
// as the result of an inline query.
//
// See https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

func (imc *InputVenueMessageContent) isInputMessageContent() {}

// InputContactMessageContent represents the content of a contact message to be
// sent as the result of an inline query.
//
// See https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

func (imc *InputContactMessageContent) isInputMessageContent() {}

// InputInvoiceMessageContent represents the content of an invoice message to be
// sent as the result of an inline query.
//
// See https://core.telegram.org/bots/api#inputinvoicemessagecontent
type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoURL                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int            `json:"photo_size,omitempty"`
	PhotoWidth                int            `json:"photo_width,omitempty"`
	PhotoHeight               int            `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

func (imc *InputInvoiceMessageContent) isInputMessageContent() {}
