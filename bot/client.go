package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// ResponseParameters contain information about why a request was unsuccessful.
type ResponseParameters struct {
	// The group has been migrated to a supergroup with the specified identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	// In case of exceeding flood control, the number of seconds to wait before making another request.
	RetryAfter int `json:"retry_after,omitempty"`
}

type tgResponse struct {
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"`
	Description string              `json:"description,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// TGError represents an error returned by the Telegram API.
// It implements the builtin error interface.
type TGError struct {
	Code        int                // HTTP-like error code returned by Telegram
	Description string             // Human-readable description of the error
	Parameters  ResponseParameters // Optional parameters for help with error handling
}

func (e *TGError) Error() string {
	return fmt.Sprintf("telegram api error: %d %s", e.Code, e.Description)
}

// IsFlood returns true if the error is a flood control error (code 429).
func (e *TGError) IsFlood() bool {
	return e.Code == 429
}

// GetRetryAfter returns the number of seconds to wait if it's a flood error.
func (e *TGError) GetRetryAfter() int {
	return e.Parameters.RetryAfter
}

// MigrateTo returns the new chat identifier if the group has been migrated.
func (e *TGError) MigrateTo() (int64, bool) {
	if e.Parameters.MigrateToChatID != 0 {
		return e.Parameters.MigrateToChatID, true
	}
	return 0, false
}

// apiRequest выполняет запрос к Telegram API.
// Если в параметрах p есть файлы, метод автоматически переключается на multipart/form-data.
// Это удобно: вам не нужно следить за тем, какой Content-Type выставить — библиотека сделает это за вас.
func (b *Bot) apiRequest(method string, p any, res any) error {
	url := b.buildURL(method)
	var contentType string
	var body io.Reader
	var err error

	// Универсальная подготовка payload.
	// extractMultipart сам разберется:
	// 1. Найдет файлы (*types.InputFile) и подготовит их.
	// 2. Сформирует map[string]any из полей структуры.
	formFields, files, packErr := extractMultipart(p)
	if packErr != nil {
		return packErr
	}

	if len(files) > 0 {
		// Если файлы есть, используем multipart/form-data
		contentType, body, err = b.prepareMultipartPayload(formFields, files)
	} else {
		// Если файлов нет, отправляем как обычный JSON.
		// Используем оригинальную структуру p, чтобы гарантировать стандартную JSON-сериализацию
		contentType, body, err = b.prepareJSONPayload(p)
	}

	if err != nil {
		b.error(fmt.Errorf("error preparing payload: %w", err))
		return err
	}
	// Выполняем запрос
	responseBody, err := b.do(url, contentType, body)
	if err != nil {
		b.error(fmt.Errorf("[http.client Error]: %w", err))
		return err
	}
	// Десериализуем ответ
	err = b.decodeResponse(responseBody, res)
	if err != nil {
		var tgErr *TGError
		if !errors.As(err, &tgErr) {
			// Если это не API-ошибка (TGError), значит произошла системная ошибка
			// (проблема с сетью на этапе чтения тела или ошибка парсинга JSON).
			// Логируем её в централизованный обработчик ошибок бота.
			b.error(fmt.Errorf("[request Error]: %w", err))
		}
	}
	return err
}

// buildURL формирует URL для запроса в Telegram API
func (b *Bot) buildURL(method string) string {
	if b.testEnv {
		return b.baseURL + "/test/" + method
	}
	return b.baseURL + "/" + method
}

// do handles the lower-level HTTP request execution and response reading.
func (b *Bot) do(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(b.ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			b.error(closeErr)
		}
	}()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Telegram sends API errors with 4xx status codes, but the body still contains
	// a valid JSON with "ok: false" and details. We return the body to let
	// decodeResponse handle it.
	return responseBody, nil
}

// decodeResponse unmarshal the raw response data into a struct.
// If the response indicates an error, it returns a *TGError.
func (b *Bot) decodeResponse(data []byte, res any) error {
	var resp tgResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("failed to decode telegram response: %w", err)
	}

	if !resp.OK {
		tgErr := &TGError{
			Code:        resp.ErrorCode,
			Description: resp.Description,
		}
		if resp.Parameters != nil {
			tgErr.Parameters = *resp.Parameters
		}
		return tgErr
	}

	if res != nil {
		if err := json.Unmarshal(resp.Result, res); err != nil {
			return fmt.Errorf("failed to decode result: %w", err)
		}
	}
	return nil
}

// prepareJSONPayload prepares a standard JSON request body.
func (b *Bot) prepareJSONPayload(params any) (string, io.Reader, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return "", nil, err
	}
	return "application/json", bytes.NewReader(data), nil
}

// prepareMultipartPayload creates a multipart form with streaming support via io.Pipe.
// We use streaming to avoid loading large files into RAM entirely.
//
//goland:noinspection GoShadowedVar
func (b *Bot) prepareMultipartPayload(formFields map[string]any, files []filePayload) (string, io.Reader, error) {
	pr, pw := io.Pipe()
	mpw := multipart.NewWriter(pw)

	go func() {
		var err error //nolint:govet
		defer func() {
			if r := recover(); r != nil {
				b.error(fmt.Errorf("panic in prepareMultipartPayload goroutine: %v", r))
				if panicErr, ok := r.(error); ok {
					err = panicErr
				} else {
					err = fmt.Errorf("%v", r)
				}
			}

			// First close the multipart writer to write the final boundary
			_ = mpw.Close()

			if err != nil {
				_ = pw.CloseWithError(err)
			} else {
				_ = pw.Close()
			}
		}()

		for key, value := range formFields {
			var valStr string
			switch v := value.(type) {
			case string:
				valStr = v
			case int, int8, int16, int32, int64:
				valStr = fmt.Sprintf("%d", v)
			case bool:
				if v {
					valStr = "true"
				}
			case nil:
				continue
			default:
				// For complex types (like ReplyMarkup) inside multipart,
				// Telegram expects a JSON string.
				data, err := json.Marshal(v)
				if err == nil {
					valStr = string(data)
				}
			}

			if valStr != "" {
				if err = mpw.WriteField(key, valStr); err != nil {
					return
				}
			}
		}

		for _, file := range files {
			err = func() error {
				if file.Reader != nil {
					part, err := mpw.CreateFormFile(file.FieldName, file.FileName)
					if err != nil {
						return err
					}
					_, err = io.Copy(part, file.Reader)
					return err
				} else if file.FilePath != "" {
					f, err := os.Open(file.FilePath)
					if err != nil {
						return err
					}

					defer func() {
						if closeErr := f.Close(); closeErr != nil {
							b.error(fmt.Errorf("failed to close file %s: %w", file.FilePath, closeErr))
						}
					}()

					part, err := mpw.CreateFormFile(file.FieldName, filepath.Base(file.FilePath))
					if err != nil {
						return err
					}
					_, err = io.Copy(part, f)
					return err
				}
				return nil
			}()

			if err != nil {
				return // Exit goroutine, defer will call pw.CloseWithError(err)
			}
		}
	}()

	return mpw.FormDataContentType(), pr, nil
}
