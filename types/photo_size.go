package types

// PhotoSizes represents a slice of PhotoSize with additional me
type PhotoSizes []PhotoSize

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
//
// See https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

// Last returns the largest photo size (the last in the slice) or nil if the
// slice is empty.
func (ps PhotoSizes) Last() *PhotoSize {
	if len(ps) == 0 {
		return nil
	}
	return &ps[len(ps)-1]
}
