package types

// File represents a file ready to be downloaded. The file can be downloaded via
// the link
//
//	https://api.telegram.org/file/bot<token>/<file_path>
//
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
//
//	The maximum file size to download is 20 MB
//
// See https://core.telegram.org/bots/api#file
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}
