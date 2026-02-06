package types

import (
	"encoding/json"
	"fmt"
)

// ChatMemberUpdated represents changes in the status of a chat member.
//
// See https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	Chat                    *Chat           `json:"chat"`
	From                    *User           `json:"from"`
	Date                    int             `json:"date"`
	OldChatMember           *ChatMemberData `json:"old_chat_member"`
	NewChatMember           *ChatMemberData `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"`
}

// EffectiveChat returns the chat where the member status was updated.
func (u *ChatMemberUpdated) EffectiveChat() *Chat {
	return u.Chat
}

// EffectiveUser returns the user who triggered the update.
func (u *ChatMemberUpdated) EffectiveUser() *User {
	return u.From
}

type ChatMemberEnum string

const (
	ChatMemberTypeOwner         ChatMemberEnum = "creator"
	ChatMemberTypeAdministrator ChatMemberEnum = "administrator"
	ChatMemberTypeMember        ChatMemberEnum = "member"
	ChatMemberTypeRestricted    ChatMemberEnum = "restricted"
	ChatMemberTypeLeft          ChatMemberEnum = "left"
	ChatMemberTypeBanned        ChatMemberEnum = "kicked"
)

// ChatMember is an interface that represents a chat member. It is implemented by
// the following types: ChatMemberOwner, ChatMemberAdministrator,
// ChatMemberMember, ChatMemberRestricted, ChatMemberLeft, ChatMemberBanned.
//
// See https://core.telegram.org/bots/api#chatmember
type ChatMember interface {
	isChatMember()
}

// ChatMemberData is a wrapper around the ChatMember interface.
//
// See https://core.telegram.org/bots/api#chatmember
type ChatMemberData struct {
	impl ChatMember
}

func (cm *ChatMemberData) MarshalJSON() ([]byte, error) {
	if cm.impl == nil {
		return []byte("null"), nil
	}
	return json.Marshal(cm.impl)
}

func (cm *ChatMemberData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Status string `json:"status"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Status {
	case string(ChatMemberTypeOwner):
		var val ChatMemberOwner
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cm.impl = &val
	case string(ChatMemberTypeAdministrator):
		var val ChatMemberAdministrator
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cm.impl = &val
	case string(ChatMemberTypeMember):
		var val ChatMemberMember
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cm.impl = &val
	case string(ChatMemberTypeRestricted):
		var val ChatMemberRestricted
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cm.impl = &val
	case string(ChatMemberTypeLeft):
		var val ChatMemberLeft
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cm.impl = &val
	case string(ChatMemberTypeBanned):
		var val ChatMemberBanned
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		cm.impl = &val
	default:
		return fmt.Errorf("unknown ChatMemberData type: %s", helper.Status)
	}

	return nil
}

func (cm *ChatMemberData) Owner() *ChatMemberOwner {
	if v, ok := cm.impl.(*ChatMemberOwner); ok {
		return v
	}
	return nil
}

func (cm *ChatMemberData) Administrator() *ChatMemberAdministrator {
	if v, ok := cm.impl.(*ChatMemberAdministrator); ok {
		return v
	}
	return nil
}

func (cm *ChatMemberData) Member() *ChatMemberMember {
	if v, ok := cm.impl.(*ChatMemberMember); ok {
		return v
	}
	return nil
}

func (cm *ChatMemberData) Restricted() *ChatMemberRestricted {
	if v, ok := cm.impl.(*ChatMemberRestricted); ok {
		return v
	}
	return nil
}

func (cm *ChatMemberData) Left() *ChatMemberLeft {
	if v, ok := cm.impl.(*ChatMemberLeft); ok {
		return v
	}
	return nil
}

func (cm *ChatMemberData) Banned() *ChatMemberBanned {
	if v, ok := cm.impl.(*ChatMemberBanned); ok {
		return v
	}
	return nil
}

// ChatMemberOwner represents a chat member that owns the chat and has all administrator privileges.
//
// See https://core.telegram.org/bots/api#chatmemberowner
type ChatMemberOwner struct {
	Status      ChatMemberEnum `json:"status"` // always “creator”
	User        *User          `json:"user"`
	IsAnonymous bool           `json:"is_anonymous"`
	CustomTitle string         `json:"custom_title,omitempty"`
}

func (cmo *ChatMemberOwner) isChatMember() {
}

// ChatMemberAdministrator represents a chat member that has some additional privileges.
//
// See https://core.telegram.org/bots/api#chatmemberadministrator
type ChatMemberAdministrator struct {
	Status                  ChatMemberEnum `json:"status"` // always “administrator”
	User                    *User          `json:"user"`
	CanBeEdited             bool           `json:"can_be_edited"`
	IsAnonymous             bool           `json:"is_anonymous"`
	CanManageChat           bool           `json:"can_manage_chat"`
	CanDeleteMessages       bool           `json:"can_delete_messages"`
	CanManageVideoChats     bool           `json:"can_manage_video_chats"`
	CanRestrictMembers      bool           `json:"can_restrict_members"`
	CanPromoteMembers       bool           `json:"can_promote_members"`
	CanChangeInfo           bool           `json:"can_change_info"`
	CanInviteUsers          bool           `json:"can_invite_users"`
	CanPostStories          bool           `json:"can_post_stories"`
	CanEditStories          bool           `json:"can_edit_stories"`
	CanDeleteStories        bool           `json:"can_delete_stories"`
	CanPostMessages         bool           `json:"can_post_messages,omitempty"`
	CanEditMessages         bool           `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool           `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool           `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool           `json:"can_manage_direct_messages,omitempty"`
	CustomTitle             string         `json:"custom_title,omitempty"`
}

func (cma *ChatMemberAdministrator) isChatMember() {}

// ChatMemberMember represents a chat member that has no additional privileges or restrictions.
//
// See https://core.telegram.org/bots/api#chatmembermember
type ChatMemberMember struct {
	Status    ChatMemberEnum `json:"status"` // always “member”
	User      *User          `json:"user"`
	UntilDate int            `json:"until_date,omitempty"`
}

func (cmm *ChatMemberMember) isChatMember() {
}

// ChatMemberRestricted represents a chat member that is under certain restrictions in the chat.
// Supergroups only.
//
// See https://core.telegram.org/bots/api#chatmemberrestricted
type ChatMemberRestricted struct {
	Status                ChatMemberEnum `json:"status"` // always “restricted”
	User                  *User          `json:"user"`
	IsMember              bool           `json:"is_member"`
	CanSendMessages       bool           `json:"can_send_messages"`
	CanSendAudios         bool           `json:"can_send_audios"`
	CanSendDocuments      bool           `json:"can_send_documents"`
	CanSendPhotos         bool           `json:"can_send_photos"`
	CanSendVideos         bool           `json:"can_send_videos"`
	CanSendVideoNotes     bool           `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool           `json:"can_send_voice_notes"`
	CanSendPolls          bool           `json:"can_send_polls"`
	CanSendOtherMessages  bool           `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool           `json:"can_add_web_page_previews"`
	CanChangeInfo         bool           `json:"can_change_info"`
	CanInviteUsers        bool           `json:"can_invite_users"`
	CanPinMessages        bool           `json:"can_pin_messages"`
	CanManageTopics       bool           `json:"can_manage_topics"`
	UntilDate             int            `json:"until_date"`
}

func (cmr *ChatMemberRestricted) isChatMember() {
}

// ChatMemberLeft represents a chat member that isn't currently a member of the
// chat, but may join it themselves.
//
// See https://core.telegram.org/bots/api#chatmemberleft
type ChatMemberLeft struct {
	Status ChatMemberEnum `json:"status"` // always “left”
	User   *User          `json:"user"`
}

func (cml *ChatMemberLeft) isChatMember() {}

// ChatMemberBanned represents a chat member that was banned in the chat
// and can't return to the chat or view chat messages.
//
// See https://core.telegram.org/bots/api#chatmemberbanned
type ChatMemberBanned struct {
	Status    ChatMemberEnum `json:"status"` // always “kicked”
	User      *User          `json:"user"`
	UntilDate int            `json:"until_date"`
}

func (cmb *ChatMemberBanned) isChatMember() {}
