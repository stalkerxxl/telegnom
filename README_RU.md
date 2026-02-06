# 🤖 Telegnom

Легкий и современный фреймворк на Go для Telegram Bot API. Вдохновленный Python-библиотеками (Aiogram, Telebot), он предоставляет Go-разработчикам удобную среду для разработки Telegram-ботов.

## ✨ Особенности

- **⚙️ Гибкие Middleware**: 3 уровня вложенности (Global, Group, Local) для полного контроля над потоком данных.
- **🔍 Мощные фильтры**: Встроенные (`IsCommand`, `Regexp`, `HasPhoto`...) и легкое создание собственных фильтров.
- **📂 Группировка хендлеров**: Логическое разделение кода и применение общих правил к наборам команд.
- **🚀 Конкурентность**: Встроенная система воркеров для параллельной и масштабируемой обработки обновлений.
- **🔗 Webhook & Polling**: Поддержка Long Polling и 3 режима Webhook (включая интеграцию в любые внешние фреймворки).
- **🛡️ Стабильность**: Автоматическое восстановление после паник (`PanicRecovery`) и типизированная обработка ошибок.
- **💎 Строгая типизация**: Никаких `any` или `interafce{}`, только строгая типизация параметров.
- **📁 Работа с файлами**: Удобная отправка медиа через `io.Reader`, путь к файлу, URL или FileID.
- **🧠 FSM (Finite State Machine)**: Удобное управление состояниями диалогов пользователя (⏳ в разработке).
- **⌨️ Keyboard Factory**: Конструктор клавиатур и типизированных Callback Data (⏳ в разработке).
- **🪶 Легкость**: Никаких внешних зависимостей — только стандартная библиотека Go.
- **📃 Документация**: Код следует стандартам GoDoc и содержит ссылки на официальную документацию Telegram Bot API.

# 📚 Содержание

- [Установка](#-установка)
- [Быстрый старт](#-быстрый-старт)
- [Опции бота](#️-конфигурация-бота)
- [Структура фреймворка](#️-структура-фреймворка)
    - [bot.Context](#-botcontext--сердце-фреймворка)
    - [bot.Bot](#️-botbot--управление)
    - [types.Update](#-typesupdate--данные)
- [Хендлеры](#-хендлеры)
- [Фильтры](#-фильтры)
- [Группы хендлеров](#-группы-хендлеров)
- [Мидлвары (Middleware)](#-мидлвары-middleware)
- [Жизненный цикл обновления](#-жизненный-цикл-обновления-lifecycle)
- [Запуск бота](#-запуск-бота)
    - [Long-Polling Mode](#-long-polling-mode)
    - [Webhook Mode](#-webhook-mode)
- [Работа с файлами](#-работа-с-файлами)
- [Обработка ошибок](#️-обработка-ошибок)
- [Логирование и отладка](#-логирование-и-отладка)
- [Управление состоянием пользователя](#-управление-состоянием-пользователя)
- [Обработка специальных типов обновлений](#обработка-специальных-типов-обновлений)
- [Работа с кнопками и клавиатурами](#работа-с-кнопками-и-клавиатурами)
- [Типичные сценарии использования](#типичные-сценарии-использования)
- [Тестирование](#тестирование)

## 📦 Установка

Убедитесь, что у вас установлен Go версии 1.25 или выше.

```bash
go get -U github.com/stalkerxxl/telegnom
```

## 🚀 Быстрый старт

Пример простого бота, который отвечает "Hello!" на команду `/start`.

```go
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/stalkerxxl/telegnom/bot"
	f "github.com/stalkerxxl/telegnom/filters"
)

func main() {
	// 1. Настройка контекста для корректного завершения (Graceful Shutdown)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 2. Инициализация бота
	b, _ := bot.NewBot(ctx, os.Getenv("BOT_TOKEN"))

	// 3. Регистрация хендлера
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Hello from Telegnom!",
		})
	}).OnMessage(f.Command("start"))

	// 4. Запуск (блокирующий вызов)
	b.StartPolling()
}
```

## ⚙️ Конфигурация бота

При создании бота с помощью функции `bot.NewBot()` можно передать различные опции для настройки его поведения. Все опции
являются необязательными.

```go
opts := []bot.Option{
bot.WithPanicRecovery(),
bot.WithNumWorkers(10),
bot.WithAllowedUpdates(types.AllUpdateTypes()...),
}

b, err := bot.NewBot(ctx, token, opts...)
```

### Доступные опции

#### Основные

- `WithApiUrl(url string)`: Использовать кастомный URL API (например, для локального Bot API сервера).
- `WithNumWorkers(n int)`: Количество параллельных воркеров для обработки обновлений (по умолчанию: 1).
- `WithClientTimeout(d time.Duration)`: Таймаут HTTP-клиента.
- `WithHTTPClient(client *http.Client)`: Использовать свой HTTP-клиент (например, для прокси).
- `WithPanicRecovery()`: Включает мидлвар для восстановления после паник в хендлерах.
- `WithTestEnv()`: Переключает бота на тестовое окружение Telegram.

#### Long Polling

- `WithAllowedUpdates(ut ...types.UpdateType)`: Список типов обновлений, которые вы хотите получать.
- `WithPollTimeout(d time.Duration)`: Тайм-аут долгого опроса (по умолчанию: 1 минута).
- `WithPollLimit(limit int)`: Максимальное количество обновлений, получаемых за один запрос (1-100).
- `WithPollOffset(offset int64)`: Смещение для получения обновлений.
- `WithUpdatesChanCap(cap int)`: Емкость внутреннего канала обновлений.

#### Webhook

- `WithWebhookURL(url string)`: Устанавливает путь для вебхука на основе переданного URL.
- `WithWebhookSecretToken(token string)`: Секретный токен для верификации запросов от Telegram.

#### Отладка и ошибки

- `WithDebugHandler(h HandlerFunc)`: Кастомный обработчик для логирования каждого входящего апдейта.
- `WithErrorHandler(h bot.ErrorHandler)`: Обработчик системных ошибок и ошибок сетевого уровня.

## 🏗️ Структура фреймворка

### 💎 bot.Context — Сердце фреймворка

`bot.Context` передается во все хендлеры, фильтры и мидлвары. Он объединяет в себе `context.Context` (Go), объект бота и
текущее обновление.

- **`tg.Bot`**: Доступ ко всем методам Telegram API (`SendMessage`, `GetMe` и т.д.).
- **`tg.Update`**: Данные текущего события.
- **`tg.ChatID() / tg.SenderID()`**: Удобные хелперы для получения ID чата/пользователя.
- **`tg.Set(key, val) / tg.Get(key)`**: Передача данных между мидлварами и хендлерами.

### 🕹️ bot.Bot — Управление

Основной объект для настройки и управления жизненным циклом бота.

- **Регистрация**: `b.Handler()`, `b.Group()`, `b.Use()`.
- **Запуск**: `b.StartPolling()`, `b.StartWebhook()`.
- **API**: Все методы Telegram Bot API в PascalCase.

### 📊 types.Update — Данные

Полная типизированная структура обновления Telegram. Содержит вспомогательные методы для быстрой проверки содержимого (
`HasText()`, `IsCommand()`, `IsCallback()`), которые автоматически находят сообщение внутри апдейта.
> Также эти встроенные фильтры вы можете использовать для создания своих кастомных фильтров.

## ⚡ Хендлеры

Хендлер в Telegnom — это конфигурируемая единица, состоящая из:

1. **Логики** (`HandlerFunc`): Что делать.
2. **Триггера** (`OnMessage`, `OnCallback`...): Когда делать.
3. **Фильтров**: Уточняющие условия (например, только команды).
4. **Локальных мидлваров**: Логика "вокруг" хендлера.

### Важные правила:

- **Порядок имеет значение**: Роутер проверяет хендлеры в том порядке, в котором они были зарегистрированы. Первый
  подошедший хендлер прерывает дальнейший поиск (First Match Wins).
- **Регистрация до запуска**: Все хендлеры и группы должны быть полностью настроены **до** вызова `StartPolling()` или
  `StartWebhook()`. После старта динамическое добавление хендлеров не поддерживается.

```go
// 1. Сначала регистрируем узкоспециализированные хендлеры
b.Handler(handleStart).OnMessage(filters.Command("start"))

// 2. Затем более общие
b.Handler(echoHandler).OnMessage(filters.HasText)

// 3. (Опционально) Хендлер с локальным мидлваром
b.Handler(secretHandler).OnMessage(filters.Command("admin")).Use(AuthMiddleware)

// 4. И только в самом конце запускаем бота
b.StartPolling()
```

> **Совет**: Если вы используете несколько триггеров (например, `.OnMessage().OnEditedMessage()`), хендлер будет
> зарегистрирован в соответствующих очередях обработки.

Подробные примеры (включая Dependency Injection и мульти-триггеры) смотрите
в [examples/handlers/README.md](examples/handlers/README.md).

### 🔍 Фильтры

Фильтры определяют, должен ли сработать данный хендлер. Они работают по принципу "все или ничего": если вы передали
несколько фильтров, все они должны вернуть `true`.

В Telegnom система фильтрации построена иерархически (от высокоуровневых к низкоуровневым):

1. **Пакет `filters`**: Готовые функции для регистрации (`f.Command("start")`, `f.Text("hello")`) и ссылки на методы
   обновления (`f.HasPhoto`, `f.IsPrivate`).
2. **Методы `types.Update`**: Логика, связанная с самим обновлением (`tg.Update.HasText()`, `tg.Update.IsCallback()`).
   Удобны тем, что сами находят нужное сообщение внутри апдейта.
3. **Методы `types.Message`**: Проверки конкретных полей сообщения (`msg.IsCommand()`, `msg.HasVideo()`).
4. **Кастомные фильтры**: Любая функция с сигнатурой `func(tg *bot.Context) bool`.

> Подробное руководство и примеры всех видов фильтров находятся
> в [examples/filters/README.md](examples/filters/README.md).

```go
// Пример использования разных уровней фильтрации
b.Handler(h1).OnMessage(filters.Command("start")) // Пакет filters
b.Handler(h2).OnMessage(func (tg *bot.Context) bool {
    return tg.Update.HasPhoto() // Методы Update
    }
)

```

## 📂 Группы хендлеров

Группы позволяют объединить несколько хендлеров и применить к ним общие мидлвары (например, для авторизации, логирования
или ограничения доступа).

```go
// Создаем группу
adminGroup := b.Group()

// Применяем мидлвар только к этой группе
adminGroup.Use(AdminCheckMiddleware)

// Все хендлеры в группе автоматически наследуют мидлвар
adminGroup.Handler(adminStats).OnMessage(filters.Command("stats"))
adminGroup.Handler(adminBan).OnMessage(filters.Command("ban"))
```

> **Важно**: Группы не могут быть вложенными. Фильтры применяются только на уровне конкретных хендлеров внутри группы.

Подробности и пример реализации защиты админ-панели смотрите в [examples/groups/README.md](examples/groups/README.md).

## 🧅 Мидлвары (Middleware)

Мидлвары позволяют выполнять код ДО и ПОСЛЕ обработки события. В Telegnom используется "луковая" архитектура:

1. **Глобальные** (`b.Use`): Для всех событий.
2. **Групповые** (`group.Use`): Для хендлеров в группе.
3. **Локальные** (`handler.Use`): Для одного конкретного хендлера.

### Принцип работы:

```go
func MyMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
    return func (tg *bot.Context) {
    // Код ДО хендлера (IN)
    next(tg)
    // Код ПОСЛЕ хендлера (OUT)
    }
}
```

> **Остановка цепочки**: Если вы не вызовете `next(tg)` и сделаете `return`, выполнение прервется, и хендлер не будет
> вызван. Это идеально подходит для авторизации.

Подробные примеры и описание передачи данных через контекст смотрите
в [examples/middleware/README.md](examples/middleware/README.md).

## 🔄 Жизненный цикл обновления (Lifecycle)

Понимание того, как Telegnom обрабатывает каждый апдейт, критически важно для написания предсказуемого кода. Весь
процесс можно представить как конвейер.

### Схема обработки

```text
[ Telegram Update ]
       ⬇
[ Global Middleware (IN) ]
       ⬇
   (Роутер ищет хендлер) <─────────┐
       ⬇                          │
[  Filters Check  ]   ─> false ────┘
       ⬇
   (Хендлер найден!)
       ⬇
[ Group Middleware (IN) ]
       ⬇
[ Local Middleware (IN) ]
       ⬇
┌───────────────────────┐
│   HANDLER FUNCTION    │
└───────────────────────┘
       ⬇
[ Local Middleware (OUT) ]
       ⬇
[ Group Middleware (OUT) ]
       ⬇
[ Global Middleware (OUT) ]
```

### Визуализация вложенности (Трассировка)

Поскольку мидлвары работают по принципу "матрешки" (или луковицы), порядок выполнения кода выглядит так. Представьте,
что у нас есть логирование на каждом уровне:

```text
[Global]  Started processing update #100
  [Group]   Checked Admin rights... OK
    [Local]   Prepared user context
      >>> HANDLER: Executing logic for /start command
    [Local]   Cleanup context
  [Group]   Saved admin stats
[Global]  Finished processing (Took 2ms)
```

### Управление потоком (Flow Control)

Вы можете остановить обработку на разных этапах, но семантика будет разной:

#### 1. Фильтры — это "Фейсконтроль"

Фильтры проверяются **до** запуска мидлваров группы и самого хендлера.

- Если фильтр вернул `false` -> Роутер просто идет искать **следующий** подходящий хендлер.
- **Сценарий**: "Эта команда доступна только в личке". Если пишут в группе — бот просто игнорирует (или ищет другой
  хендлер).

#### 2. Мидлвары — это "Охрана"

Мидлвар запускается, когда хендлер уже найден.

- Если мидлвар не вызвал `next(tg)` -> Обработка **прекращается** полностью. Никакие другие хендлеры больше не ищутся.
- **Сценарий**: "Команда найдена, но у юзера нет прав". Мидлвар пишет "Доступ запрещен" и делает `return`.

## 🏁 Запуск бота

Telegnom поддерживает два способа получения обновлений: **Long Polling** и **Webhook**.
[Документация Telegram](https://core.telegram.org/bots/api#making-requests).

- **Long Polling**: Бот сам периодически опрашивает серверы Telegram на наличие новых событий. Это самый простой способ,
  не требующий внешнего IP-адреса или SSL-сертификата. Идеально подходит для локальной разработки и небольших ботов.
- **Webhook**: Telegram сам отправляет HTTP-запрос на ваш сервер, как только происходит событие. Этот метод более
  эффективен и быстрее реагирует на сообщения, но требует наличия публичного URL и настроенного HTTPS. Рекомендуется для
  высоконагруженных проектов.

### 📡 Long-Polling Mode

Метод `StartPolling()` запускает бесконечный цикл опроса Telegram API на наличие обновлений. Он является блокирующим —
выполнение программы останавливается до прерывания (например, Ctrl+C).

Опции (необязательные) Long Polling ([TG API DOC](https://core.telegram.org/bots/api#getupdates)):

- `WithPollLimit(limit int)`: Устанавливает максимальное количество обновлений, получаемых за один запрос (от 1 до 100).
- `WithPollOffset(offset int64)`: Устанавливает начальный ID обновления для polling. Положительное значение — начать с
  конкретного ID; отрицательное (например, -1) — получить только последнее обновление.
- `WithPollTimeout(d time.Duration)`: Устанавливает тайм-аут для одного запроса polling.
- `WithAllowedUpdates(ut ...types.UpdateType)`: Устанавливает типы обновлений, которые бот будет получать.

```go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/stalkerxxl/telegnom/bot"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	token := os.Getenv("BOT_TOKEN")

	// 2. Инициализация бота
	b, err := bot.NewBot(ctx, token, bot.WithPanicRecovery())
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	// 3. Регистрация хендлеров
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Работаю в режиме Long Polling!",
		})
	}).OnMessage()

	// 4. Запуск бесконечного цикла опроса
	log.Println("Бот запущен в режиме Long Polling...")
	b.StartPolling()
}

```

> `bot.StartPolling()` вернет ошибку, если настроен исходящий вебхук. Рекомендуем вызывать `bot.DeleteWebhook()` перед
> запуском.

### 🔗 Webhook Mode

В режиме Webhook Telegram сам отправляет обновления на ваш сервер. Это исключает задержки и экономит ресурсы. Telegnom
предлагает три способа работы с вебхуками:

1. **Встроенный HTTP-сервер**: `bot.StartWebhook(":8080")` — самый простой способ (обычно используется за Nginx |
   Ngrok).
2. **Встроенный HTTPS-сервер**: `bot.StartWebhookTLS(":443", "cert.pem", "key.pem")` — если бот смотрит напрямую в
   интернет.
3. **Внешний сервер**: `bot.WebhookHandler()` — для интеграции в существующее веб-приложение (Gin, Echo, и т.д.).

> **Важно**: Перед запуском бота в этом режиме вы должны самостоятельно вызвать `tg.Bot.SetWebhook()` (один раз при
> смене режима или изменении параметров)

#### Пример использования (встроенный сервер)

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/stalkerxxl/telegnom/bot"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("BOT_TOKEN")

	// 1. Инициализация бота с настройкой вебхука
	webhookURL := "https://my-domain.com/secret-path"
	secretToken := "my-secret-token"

	b, _ := bot.NewBot(ctx, token,
		bot.WithWebhookURL(webhookURL),
		bot.WithWebhookSecretToken(secretToken),
	)

	// 2. Регистрация вебхука в Telegram (делается один раз или при смене настроек)
	_, err = b.SetWebhook(&bot.SetWebhookParams{
		URL:         webhookURL,
		SecretToken: secretToken,
	})
	if err != nil {
		log.Fatalf("Ошибка регистрации вебхука: %v", err)
	}

	// 3. Регистрация хендлеров
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Получено через Webhook!",
		})
	}).OnMessage()

	// 3. Запуск встроенного сервера
	log.Println("Бот запущен в режиме Webhook на порту :8080...")
	if err := b.StartWebhook(":8080"); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
```

## 📁 Работа с файлами

Для отправки файлов (фото, видео, документы) используется структура `types.InputFile`. Вы можете отправлять файлы по
URL, FileID, пути на диске или из `io.Reader`.

Подробная документация: [docs/InputFile.md](docs/InputFile.md).

```go
// Отправка фото с диска
_, _ = tg.Bot.SendPhoto(&bot.SendPhotoParams{
    ChatID: tg.ChatID(),
    Photo:  &types.InputFile{FilePath: "images/cat.jpg"},
    Caption: "Котик",
})
```

## ⚠️ Обработка ошибок

В Telegnom обработка ошибок разделена на два уровня: **системные ошибки** и **ошибки Telegram API**. Это сделано для
того, чтобы вы могли гибко реагировать на проблемы с сетью отдельно от логических ошибок бота.

### Типы ошибок

1. **Системные ошибки**: Проблемы с интернет-соединением, DNS, ошибки парсинга некорректного JSON. Эти ошибки
   возвращаются как стандартные ошибки Go.
2. **Ошибки Telegram API**: Ошибки, которые возвращает сам Telegram (например, "Chat not found", "Bot was blocked by the
   user", "Flood control exceeded"). Они возвращаются как тип `*bot.TGError`.

### Как обрабатывать ошибки API

В Go для проверки типа ошибки используется функция `errors.As`.

```go
package main

import (
	"errors"
	"fmt"
	_ "time"

	"github.com/stalkerxxl/telegnom/bot"
)

func myHandler(tg *bot.Context) {
	_, err := tg.Bot.SendMessage(&bot.SendMessageParams{
		ChatID: tg.Update.Chat.ID,
		Text:   "Hello!",
	})

	if err != nil {
		var tgErr *bot.TGError
		if errors.As(err, &tgErr) {
			// Это ошибка от самого Telegram
			fmt.Printf("API Error: %d - %s\n", tgErr.Code, tgErr.Description)

			// Специальные методы для удобства:
			if tgErr.IsFlood() {
				retryAfter := tgErr.GetRetryAfter()
				fmt.Printf("Нужно подождать %d секунд\n", retryAfter)
			}

			if newChatID, ok := tgErr.MigrateTo(); ok {
				fmt.Printf("Чат переехал в супергруппу с ID: %d\n", newChatID)
			}
		} else {
			// Это системная ошибка (сеть, прокси и т.д.)
			fmt.Printf("System Error: %v\n", err)
		}
	}
}
```

### Централизованный обработчик (ErrorHandler)

По умолчанию в библиотеке используется встроенный обработчик, который логирует системные ошибки в консоль через
`log.Printf`. Вы можете заменить его на свой или полностью отключить.

- **По умолчанию:** Включен (встроенный логгер).
- **Свой обработчик:** Передайте функцию в `bot.WithErrorHandler(myFunc)`.
- **Отключение:** Передайте `nil` в `bot.WithErrorHandler(nil)`.

Использование кастомного обработчика полезно для отправки отчетов в Sentry, записи в файлы или уведомления админа в
Telegram.

```go
opts := []bot.Option{
// Свой обработчик
bot.WithErrorHandler(func (err error) {
    log.Printf("Критическая ошибка: %v", err)
    }),
}
// Или отключение (логи не будут выводиться)
// bot.WithErrorHandler(nil),
b, _ := bot.NewBot(ctx, token, opts...)
```

> **Важно**: Ошибки API (`TGError`) **не попадают** в `ErrorHandler` автоматически при вызовах методов вроде
`SendMessage`. Они возвращаются вам напрямую для ручной обработки. В `ErrorHandler` попадают ошибки сетевого уровня и
> ошибки, возникшие внутри самого фреймворка (например, при получении обновлений).

## 📝 Логирование и отладка

Для упрощения разработки и мониторинга в Telegnom встроены инструменты отладки.

### Просмотр входящих обновлений (DebugHandler)

Если ваш хендлер не срабатывает, скорее всего, дело в фильтрах. Чтобы увидеть, какие именно данные приходят от Telegram,
используйте `WithDebugHandler`.

```go
opts := []bot.Option{
// Использовать встроенный логгер (выводит Update в JSON)
bot.WithDebugHandler(nil), // nil выключает стандартный обработчик

// Или написать свой
bot.WithDebugHandler(func (tg *bot.Context) {
    log.Printf("ID: %d, Text: %s", tg.Update.ID, tg.Update.Message.Text)
    }),
}
```

### Восстановление после паник (PanicRecovery)

Чтобы одна ошибка в коде хендлера не уронила всего бота, всегда используйте `WithPanicRecovery()`. Она перехватывает
паники, логирует их через `ErrorHandler` и позволяет боту продолжить работу.

```go
b, _ := bot.NewBot(ctx, token, bot.WithPanicRecovery())
```

## 👤 Управление состоянием пользователя

Coming Soon...

## 📂 Структура проекта

Краткий обзор организации исходного кода Telegnom:

- **`bot/`**: 🧠 Ядро фреймворка. Содержит логику работы бота, управление контекстом и роутинг.
    - **`methods.go`**: 🛠️ Реализация всех методов Telegram Bot API (`SendMessage`, `GetMe` и др.).
    - **`methods_params.go`**: 📝 Структуры параметров для методов API (например, `SendMessageParams`).
- **`types/`**: 📊 Типы данных. Полный набор структур, соответствующих объектам Telegram Bot API (Update, Message, User и
  др.).
- **`filters/`**: 🔍 Пакет со встроенными фильтрами для гибкой настройки триггеров хендлеров.
- **`examples/`**: 🎓 Примеры использования. Готовые рецепты по работе с мидлварами, группами, вебхуками и многим другим.

## Обработка специальных типов обновлений

Coming Soon...

## Работа с кнопками и клавиатурами

Coming Soon...

## Типичные сценарии использования

Coming Soon...

## Тестирование

Coming Soon...
