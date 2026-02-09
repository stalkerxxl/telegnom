package bot

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/stalkerxxl/telegnom/types"
)

// filePayload for transferring files in multipart requests.
type filePayload struct {
	FieldName string    // Name of the field in the form (e.g., "photo", "document")
	FilePath  string    // Path to the file (optional)
	Reader    io.Reader // Or a ready reader (optional)
	FileName  string    // File name (required for Reader)
}

// extractMultipart (it's a MAGIC 😁) takes a struct with potential *InputFile fields and extracts:
// 1. A map of field names to values for the API request.
// 2. A slice of filePayloads for any files that need to be uploaded via multipart/form-data.
//
// It handles both top-level *InputFile fields (using the "media" tag) and nested *InputFile fields.
// For nested files, it assigns unique AttachName values and prepares them for upload.
func extractMultipart(p any) (map[string]any, []filePayload, error) {
	fields := make(map[string]any)
	var files []filePayload
	fileCounter := 0

	v := reflect.ValueOf(p)

	// If nil or a non-structural type is passed, we return empty fields and files without error.
	if !v.IsValid() {
		return fields, files, nil
	}

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return fields, files, nil
		}
		v = v.Elem()
	}

	// We expect a structure for parsing tags. If this is not a structure, return nil
	if v.Kind() != reflect.Struct {
		return fields, files, nil
	}

	t := v.Type()

	// 1. First, recursively search for all *InputFiles that need to be loaded.
	// It is important to do this before forming the fields so that the files already have an AttachName.
	var findFiles func(rv reflect.Value)
	findFiles = func(rv reflect.Value) {
		if !rv.IsValid() {
			return
		}

		// If we can't access the interface (the field is private),
		// just skip it to avoid panic.
		if !rv.CanInterface() {
			return
		}

		switch rv.Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				return
			}
			// Check if the pointer is *InputFile itself
			if inputFile, ok := rv.Interface().(*types.InputFile); ok {
				if inputFile != nil && inputFile.ID == "" && inputFile.URL == "" && (inputFile.Path != "" || inputFile.Reader != nil) {
					// If this is a local file or Reader, and it does not yet have an attachment name
					if inputFile.AttachName == "" {
						attachName := fmt.Sprintf("file_%d", fileCounter)
						fileCounter++
						inputFile.AttachName = attachName
						files = append(files, filePayload{
							FieldName: attachName,
							FilePath:  inputFile.Path,
							Reader:    inputFile.Reader,
							FileName:  inputFile.Name,
						})
					}
				}
				return
			}
			findFiles(rv.Elem())
		case reflect.Slice, reflect.Array:
			for i := 0; i < rv.Len(); i++ {
				findFiles(rv.Index(i))
			}
		case reflect.Struct:
			for i := 0; i < rv.NumField(); i++ {
				// Check if the field is exportable
				if rv.Type().Field(i).PkgPath != "" {
					continue
				}
				findFiles(rv.Field(i))
			}
		case reflect.Interface:
			if !rv.IsNil() {
				findFiles(rv.Elem())
			}
		default:
			return
		}
	}

	findFiles(v)

	// 2. Create a map of fields for the request
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Process files via the "media" tag (top level)
		mediaTag := fieldType.Tag.Get("media")
		if mediaTag != "" {
			if field.IsNil() {
				continue
			}
			inputFile, ok := field.Interface().(*types.InputFile)
			if !ok {
				continue
			}

			if err := inputFile.Verify(); err != nil {
				return nil, nil, fmt.Errorf("error verifying file for field %s: %w", fieldType.Name, err)
			}

			if inputFile.ID != "" {
				fields[mediaTag] = inputFile.ID
			} else if inputFile.URL != "" {
				fields[mediaTag] = inputFile.URL
			} else if inputFile.AttachName != "" {
				fields[mediaTag] = "attach://" + inputFile.AttachName
			}
			continue
		}

		// Processing text fields via json tag
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		tagParts := strings.Split(jsonTag, ",")
		fieldName := tagParts[0]

		if field.IsZero() && strings.Contains(jsonTag, "omitempty") {
			continue
		}

		fields[fieldName] = field.Interface()
	}

	return fields, files, nil
}
