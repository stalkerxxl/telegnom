package types

// User represents a Telegram user or bot.
//
// See https://core.telegram.org/bots/api#user
type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportInlineQueries    bool   `json:"support_inline_queries,omitempty"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           bool   `json:"has_main_web_app,omitempty"`
	HasTopicsEnabled        bool   `json:"has_topics_enabled,omitempty"`
}

// UserProfilePhotos represent a user's profile pictures.
//
// See https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int          `json:"total_count"`
	Photos     []PhotoSizes `json:"photos"`
}

// UserRating describes the rating of a user based on their Telegram Star spending's.
//
// See https://core.telegram.org/bots/api#userrating
type UserRating struct {
	Level              int `json:"level"`
	Rating             int `json:"rating"`
	CurrentLevelRating int `json:"current_level_rating"`
	NextLevelRating    int `json:"next_level_rating,omitempty"`
}
