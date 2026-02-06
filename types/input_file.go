package types

import (
	"encoding/json"
	"fmt"
	"io"
)

// InputFile is a universal structure for upload file to Telegram.
// Supports ID, local path, URL, or io.Reader.
type InputFile struct {
	ID     string
	Path   string
	URL    string
	Reader io.Reader
	Name   string

	AttachName string `json:"-"` // for attach://
}

// MarshalJSON implements the json.Marshaler interface.
// Telegram API expects in media fields either a string (file_id/URL) or "attach://name".
func (f *InputFile) MarshalJSON() ([]byte, error) {
	if f == nil {
		return []byte("null"), nil
	}

	if f.ID != "" {
		return json.Marshal(f.ID)
	}

	if f.URL != "" {
		return json.Marshal(f.URL)
	}

	if f.AttachName != "" {
		return json.Marshal("attach://" + f.AttachName)
	}

	return []byte("null"), nil
}

// Verify checks that at least one file sending method is filled.
// This will protect us from sending an empty request to Telegram.
func (f *InputFile) Verify() error {
	if f == nil {
		return nil
	}
	if f.ID == "" && f.Path == "" && f.URL == "" && f.Reader == nil {
		return fmt.Errorf("InputFile must have at least one source: ID, Path, URL, or Reader")
	}
	return nil
}
