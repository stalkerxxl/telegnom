package bot

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/stalkerxxl/telegnom/types"
)

// filePayload для передачи файлов в multipart-запросах.
type filePayload struct {
	FieldName string    // Имя поля в форме (например, "photo", "document")
	FilePath  string    // Путь к файлу (опционально)
	Reader    io.Reader // Или готовый reader (опционально)
	FileName  string    // Имя файла (обязательно для Reader)
}

// extractMultipart – универсальный помощник, который заменяет долгую ручную
// сборку полей. Для Python-разработчика это аналог того, как если бы вы
// использовали inspect или __dict__ для автоматического формирования словаря для
// requests.post(..., files=files, data=data). В Go это делается через reflection
// (пакет reflect).
func extractMultipart(p any) (map[string]any, []filePayload, error) {
	fields := make(map[string]any)
	var files []filePayload
	fileCounter := 0

	v := reflect.ValueOf(p)

	// Если передан nil (например, GetMe), просто возвращаем пустой результат
	if !v.IsValid() {
		return fields, files, nil
	}

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return fields, files, nil
		}
		v = v.Elem()
	}

	// Мы ожидаем структуру для разбора тегов. Если это не структура, возвращаем пусто.
	if v.Kind() != reflect.Struct {
		return fields, files, nil
	}

	t := v.Type()

	// 1. Сначала рекурсивно ищем все *InputFile, которые нужно загрузить.
	// Это важно сделать до формирования полей, чтобы у файлов уже были AttachName.
	var findFiles func(rv reflect.Value)
	findFiles = func(rv reflect.Value) {
		if !rv.IsValid() {
			return
		}

		// Если мы не можем получить доступ к интерфейсу (поле приватное),
		// просто пропускаем его, чтобы избежать паники.
		if !rv.CanInterface() {
			return
		}

		switch rv.Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				return
			}
			// Проверяем, не является ли указатель самим *InputFile
			if inputFile, ok := rv.Interface().(*types.InputFile); ok {
				if inputFile != nil && inputFile.ID == "" && inputFile.URL == "" && (inputFile.Path != "" || inputFile.Reader != nil) {
					// Если это локальный файл или Reader, и у него еще нет имени аттача
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
				// Проверяем, экспортируемо ли поле (начинается ли с большой буквы)
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

	// 2. Формируем мапу полей для запроса
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Обработка файлов через тег "media" (верхний уровень)
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

		// Обработка текстовых полей через тег json
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
