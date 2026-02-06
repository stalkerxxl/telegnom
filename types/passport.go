package types

import (
	"encoding/json"
	"fmt"
)

// Telegram Passport is a unified authorization method for services that require
// personal identification. Users can upload their documents once, then instantly
// share their data with services that require real-world ID (finance, ICOs,
// etc.).
//
// Please see https://core.telegram.org/passport for details.

// PassportData describes Telegram Passport data shared with the bot by the user.
//
// See https://core.telegram.org/bots/api#passportdata
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials      `json:"credentials"`
}

// PassportFile represents a file uploaded to Telegram Passport. Currently, all
// Telegram Passport files are in JPEG format when decrypted and don't exceed
// 10MB.
//
// See https://core.telegram.org/bots/api#passportfile
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

// EncryptedPassportElement describes documents or other Telegram Passport
// elements shared with the bot by the user.
//
// See https://core.telegram.org/bots/api#encryptedpassportelement
type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

// EncryptedCredentials describes data required for decrypting and authenticating
// EncryptedPassportElement. See the Telegram Passport Documentation for a
// complete description of the data decryption and authentication processes
//
// See https://core.telegram.org/bots/api#encryptedcredentials
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// PassportElementError is an interface for different types of passport element
// errors. It can be one of the following types: PassportElementErrorDataField ||
// PassportElementErrorFrontSide || PassportElementErrorReverseSide ||
// PassportElementErrorSelfie || PassportElementErrorFile ||
// PassportElementErrorFiles || PassportElementErrorTranslationFile ||
// PassportElementErrorTranslationFiles || PassportElementErrorUnspecified
//
// See https://core.telegram.org/bots/api#passportelementerror
type PassportElementError interface {
	isPassportElementError()
}

// PassportElementErrorData is a wrapper for the different types of passport element errors.
//
// See https://core.telegram.org/bots/api#passportelementerror
type PassportElementErrorData struct {
	impl PassportElementError
}

func (pee *PassportElementErrorData) UnmarshalJSON(b []byte) error {
	var helper struct {
		Source string `json:"source"`
	}

	if err := json.Unmarshal(b, &helper); err != nil {
		return err
	}

	switch helper.Source {
	case "data":
		var val PassportElementErrorDataField
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	case "front_side":
		var val PassportElementErrorFrontSide
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	case "reverse_side":
		var val PassportElementErrorReverseSide
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	case "selfie":
		var val PassportElementErrorSelfie
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	case "file":
		var val PassportElementErrorFile
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	case "files":
		var val PassportElementErrorFiles
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	case "translation_file":
		var val PassportElementErrorTranslationFile
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	case "translation_files":
		var val PassportElementErrorTranslationFiles
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	case "unspecified":
		var val PassportElementErrorUnspecified
		if err := json.Unmarshal(b, &val); err != nil {
			return err
		}
		pee.impl = &val
	default:
		return fmt.Errorf("unsupported PassportElementError type, %v", helper.Source)
	}

	return nil
}

func (pee *PassportElementErrorData) DataField() *PassportElementErrorDataField {
	if v, ok := pee.impl.(*PassportElementErrorDataField); ok {
		return v
	}
	return nil
}

func (pee *PassportElementErrorData) FrontSide() *PassportElementErrorFrontSide {
	if v, ok := pee.impl.(*PassportElementErrorFrontSide); ok {
		return v
	}
	return nil
}

func (pee *PassportElementErrorData) ReverseSide() *PassportElementErrorReverseSide {
	if v, ok := pee.impl.(*PassportElementErrorReverseSide); ok {
		return v
	}
	return nil
}

func (pee *PassportElementErrorData) Selfie() *PassportElementErrorSelfie {
	if v, ok := pee.impl.(*PassportElementErrorSelfie); ok {
		return v
	}
	return nil
}

func (pee *PassportElementErrorData) File() *PassportElementErrorFile {
	if v, ok := pee.impl.(*PassportElementErrorFile); ok {
		return v
	}
	return nil
}

func (pee *PassportElementErrorData) Files() *PassportElementErrorFiles {
	if v, ok := pee.impl.(*PassportElementErrorFiles); ok {
		return v
	}
	return nil
}

func (pee *PassportElementErrorData) TranslationFile() *PassportElementErrorTranslationFile {
	if v, ok := pee.impl.(*PassportElementErrorTranslationFile); ok {
		return v
	}
	return nil
}

func (pee *PassportElementErrorData) TranslationFiles() *PassportElementErrorTranslationFiles {
	if v, ok := pee.impl.(*PassportElementErrorTranslationFiles); ok {
		return v
	}
	return nil
}

func (pee *PassportElementErrorData) Unspecified() *PassportElementErrorUnspecified {
	if v, ok := pee.impl.(*PassportElementErrorUnspecified); ok {
		return v
	}
	return nil
}

// PassportElementErrorDataField represents an issue in one of the data fields
// that was provided by the user. The error is considered resolved when the
// field's value changes.
//
// See https://core.telegram.org/bots/api#passportelementerrordatafield
type PassportElementErrorDataField struct {
	Source    string `json:"source"` // must be "data"
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

func (pee *PassportElementErrorDataField) isPassportElementError() {}

// PassportElementErrorFrontSide represents an issue with the front side of a
// document. The error is considered resolved when the file with the front side
// of the document changes.
//
// See https://core.telegram.org/bots/api#passportelementerrorfrontside
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"` // must be "front_side"
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (pee *PassportElementErrorFrontSide) isPassportElementError() {}

// PassportElementErrorReverseSide represents an issue with the reverse side of a
// document. The error is considered resolved when the file with reverse side of
// the document changes.
//
// See https://core.telegram.org/bots/api#passportelementerrorreverseside
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"` // must be "reverse_side"
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (pee *PassportElementErrorReverseSide) isPassportElementError() {}

// PassportElementErrorSelfie represents an issue with the selfie with a
// document. The error is considered resolved when the file with the selfie
// changes.
//
// See https://core.telegram.org/bots/api#passportelementerrorselfie
type PassportElementErrorSelfie struct {
	Source   string `json:"source"` // must be "selfie"
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (pee *PassportElementErrorSelfie) isPassportElementError() {}

// PassportElementErrorFile represents an issue with a document scan. The error
// is considered resolved when the file with the document scan changes.
//
// See https://core.telegram.org/bots/api#passportelementerrorfile
type PassportElementErrorFile struct {
	Source   string `json:"source"` // must be "file"
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (pee *PassportElementErrorFile) isPassportElementError() {}

// PassportElementErrorFiles represents an issue with a list of scans. The error
// is considered resolved when the list of files containing the scans changes.
//
// See https://core.telegram.org/bots/api#passportelementerrorfiles
type PassportElementErrorFiles struct {
	Source     string   `json:"source"` // must be "files"
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (pee *PassportElementErrorFiles) isPassportElementError() {}

// PassportElementErrorTranslationFile represents an issue with one of the files
// that constitute the translation of a document. The error is considered
// resolved when the file changes.
//
// See https://core.telegram.org/bots/api#passportelementerrortranslationfile
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"` // must be "translation_file"
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (pee *PassportElementErrorTranslationFile) isPassportElementError() {}

// PassportElementErrorTranslationFiles represents an issue with the translated
// version of a document. The error is considered resolved when a file with the
// document translation change.
//
// See https://core.telegram.org/bots/api#passportelementerrortranslationfiles
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"` // must be "translation_files"
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (pee *PassportElementErrorTranslationFiles) isPassportElementError() {}

// PassportElementErrorUnspecified represents an issue in an unspecified place.
// The error is considered resolved when new data is added.
//
// See https://core.telegram.org/bots/api#passportelementerrorunspecified
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"` // must be "unspecified"
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

func (pee *PassportElementErrorUnspecified) isPassportElementError() {}
