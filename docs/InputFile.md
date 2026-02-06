# Работа с файлами в Telegnom

Библиотека `telegnom` предоставляет универсальный и прозрачный механизм для отправки файлов. Вам не нужно думать о `multipart/form-data`, границах (boundaries) или ручном формировании HTTP-запросов. Все это происходит автоматически.

Центральным элементом является структура `types.InputFile`.

## Структура InputFile

Структура `types.InputFile` позволяет указать источник файла одним из четырех способов:

```go
type InputFile struct {
    FileID   string    // Если файл уже загружен в Telegram
    URL      string    // Прямая ссылка на файл (Telegram скачает его сам)
    FilePath string    // Путь к файлу на локальном диске
    Reader   io.Reader // Любой поток данных (например, созданный в памяти)
    FileName string    // Имя файла (обязательно при использовании Reader)
}
```

Библиотека автоматически определяет, как отправить запрос:
1. Если указан `FileID` или `URL`, отправляется обычный JSON-запрос.
2. Если указан `FilePath` или `Reader`, библиотека переключается в режим `multipart/form-data`, загружает файл и автоматически проставляет ссылки `attach://`.

---

## Примеры использования

### 1. Отправка локального файла (по пути)

Самый частый сценарий. Достаточно указать путь к файлу.

```go
params := bot.SendPhotoParams{
    ChatID: 123456789,
    Photo: &types.InputFile{
        FilePath: "./images/cat.jpg",
    },
    Caption: "Смотри, какой кот!",
}

resp, err := b.SendPhoto(ctx, params)
```

### 2. Отправка по FileID

Используется, если файл уже есть на серверах Telegram (например, вы сохранили его ID из предыдущего сообщения). Это самый быстрый способ, так как не требует повторной загрузки данных.

```go
params := bot.SendDocumentParams{
    ChatID: 123456789,
    Document: &types.InputFile{
        FileID: "BQACAgIAAxkDAAI...", 
    },
}
```

### 3. Отправка по URL

Вы можете передать прямую ссылку на файл. Telegram сам скачает его и отправит пользователю.

```go
params := bot.SendPhotoParams{
    ChatID: 123456789,
    Photo: &types.InputFile{
        URL: "https://example.com/image.png",
    },
}
```

### 4. Отправка из памяти (io.Reader)

Полезно, если вы генерируете файл "на лету" (например, PDF-отчет, картинку или лог), и не хотите сохранять его на диск.
**Важно:** При использовании `Reader` обязательно указывайте `FileName`.

```go
// Пример: создаем текстовый файл прямо в памяти
content := bytes.NewBufferString("Hello, World!")

params := bot.SendDocumentParams{
    ChatID: 123456789,
    Document: &types.InputFile{
        Reader:   content,
        FileName: "hello.txt", // Имя файла, которое увидит пользователь
    },
}

resp, err := b.SendDocument(ctx, params)
```

---

## Отправка альбомов (Media Group)

Метод `SendMediaGroup` позволяет отправить несколько файлов (фото или видео) одним сообщением (альбомом).
Для этого используется слайс интерфейсов `types.InputMediaOnly`.

Библиотека рекурсивно обходит этот слайс, находит все файлы, которые нужно загрузить, формирует `multipart` запрос и корректно расставляет `attach://` ссылки для каждого элемента альбома.

```go
// Подготавливаем группу медиа
mediaGroup := []types.InputMediaOnly{
    // Фото 1: Локальный файл
    &types.InputMediaPhoto{
        Media: &types.InputFile{
            FilePath: "./photos/party_1.jpg",
        },
        Caption: "Начало вечеринки",
    },
    // Фото 2: Файл из интернета
    &types.InputMediaPhoto{
        Media: &types.InputFile{
            URL: "https://example.com/party_2.jpg",
        },
    },
    // Видео: По FileID
    &types.InputMediaVideo{
        Media: &types.InputFile{
            FileID: "BAACAgIAAx...",
        },
    },
}

params := bot.SendMediaGroupParams{
    ChatID: 123456789,
    Media:  mediaGroup,
}

resp, err := b.SendMediaGroup(ctx, params)
```

## Как это работает внутри?

Когда вы вызываете метод API (например, `SendPhoto`), происходит следующее:

1. Метод `extractMultipart` (внутри `client.go`) через рефлексию сканирует структуру параметров.
2. Он ищет все поля типа `*types.InputFile`.
3. Если находится файл, который нужно загрузить (`FilePath` или `Reader`), ему присваивается уникальное имя вложения (например, `file_0`).
4. В JSON-поле параметра подставляется строка `attach://file_0`.
5. Реальный поток данных файла добавляется в multipart-запрос под именем `file_0`.
6. Если файлов нет, структура отправляется как обычный JSON.
