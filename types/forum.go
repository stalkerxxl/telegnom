package types

// ForumTopicCreated represents a service message about a new forum topic created
// in the chat.
//
// See https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
}

// ForumTopicClosed represents a service message about a forum topic closed in
// the chat. Currently, holds no information.
//
// See https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct{}

// ForumTopicEdited represents a service message about an edited forum topic.
//
// See https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicReopened represents a service message about a forum topic reopened
// in the chat. Currently, holds no information.
//
// See https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct {
}

// GeneralForumTopicHidden represents a service message about General forum topic
// hidden in the chat. Currently, holds no information.
//
// See https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct {
}

// GeneralForumTopicUnhidden represents a service message about General forum
// topic unhidden in the chat. Currently, holds no information.
//
// See https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct {
}

// ForumTopic represents a forum topic.
//
// See https://core.telegram.org/bots/api#forumtopic
type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
}
