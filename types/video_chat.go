package types

// VideoChatScheduled represents a service message about a video chat scheduled
// in the chat.
//
// See https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

// VideoChatStarted represents a service message about a video chat started in
// the chat. Currently, holds no information.
//
// See https://core.telegram.org/bots/api#videochatstarted
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a video chat ended in the
// chat.
//
// See https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited represents a service message about new members
// invited to a video chat.
//
// See https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}
