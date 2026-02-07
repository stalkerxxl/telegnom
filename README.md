# 🤖 Telegnom

[![Go Report Card](https://goreportcard.com/report/github.com/stalkerxxl/telegnom)](https://goreportcard.com/report/github.com/stalkerxxl/telegnom)

> Support Telegram Bot API: [v9.3](https://core.telegram.org/bots/api#december-31-2025) - December 31, 2025 | [Telegram Group](https://t.me/telegnom_chat)

A lightweight and idiomatic Go-framework for the Telegram Bot API. Inspired by Python libraries (Aiogram, Telebot), it provides Go developers with a convenient environment for developing Telegram bots.

## ✨ Features

- **⚙️ Flexible Middleware**: 3 levels of nesting (Global, Group, Local) for full control over the data flow.
- **🔍 Powerful Filters**: Built-in (`IsCommand`, `Regexp`, `HasPhoto`...) and easy creation of custom filters.
- **📂 Handler Grouping**: Logical separation of code and application of common rules to sets of commands.
- **🚀 Concurrency**: Built-in worker system for parallel and scalable update processing.
- **🔗 Webhook & Polling**: Support for Long Polling and 3 Webhook modes (including integration into any external frameworks).
- **🛡️ Stability**: Automatic panic recovery (`PanicRecovery`) and typed error handling.
- **💎 Strict Typing**: No `any` or `interface{}`, only strict parameter typing.
- **📁 File Handling**: Convenient media sending via `io.Reader`, file path, URL, or FileID.
- **🧠 FSM (Finite State Machine)**: Convenient user dialog state management (⏳ in development).
- **⌨️ Keyboard Factory**: Keyboard constructor and typed Callback Data (⏳ in development).
- **🪶 Lightweight**: No external dependencies — only the standard Go library.
- **📃 Documentation**: Code follows GoDoc standards and contains links to the official Telegram Bot API documentation.

## 📚 Content

- [Installation](#-installation)
- [Quick Start](#-quick-start)
- [Bot Configuration](#️-bot-configuration)
- [Framework Structure](#️-framework-structure)
    - [bot.Context](#-botcontext--heart-of-the-framework)
    - [bot.Bot](#️-botbot--control)
    - [types.Update](#-typesupdate--data)
- [Handlers](#-handlers)
- [Filters](#-filters)
- [Handler Groups](#-handler-groups)
- [Middleware](#-middleware)
- [Update Lifecycle](#-update-lifecycle)
- [Running the Bot](#-running-the-bot)
    - [Long-Polling Mode](#-long-polling-mode)
    - [Webhook Mode](#-webhook-mode)
- [File Handling](#-file-handling)
- [Error Handling](#️-error-handling)
- [Logging and Debugging](#-logging-and-debugging)
- [User State Management](#-user-state-management)
- [Handling Special Update Types](#handling-special-update-types)
- [Working with Keyboards](#working-with-keyboards)
- [Typical Usage Scenarios](#typical-usage-scenarios)
- [Testing](#testing)

## 📦 Installation

Ensure you have Go version 1.25 or higher installed.

```bash
go get -U github.com/stalkerxxl/telegnom
```

## 🚀 Quick Start

Example of a simple bot that replies "Hello!" to the `/start` command.

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
	// 1. Setup context for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 2. Initialize the bot
	b, _ := bot.NewBot(ctx, os.Getenv("BOT_TOKEN"))

	// 3. Register a handler
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Hello from Telegnom!",
		})
	}).OnMessage(f.Command("start"))

	// 4. Start (blocking call)
	b.StartPolling()
}
```

## ⚙️ Bot Configuration

When creating a bot using `bot.NewBot()`, you can pass various options to configure its behavior. All options are optional.

```go
opts := []bot.Option{
    bot.WithPanicRecovery(),
    bot.WithNumWorkers(10),
    bot.WithAllowedUpdates(types.AllUpdateTypes()...),
}

b, err := bot.NewBot(ctx, token, opts...)
```

### Available Options

#### Basic

- `WithApiUrl(url string)`: Use a custom API URL (e.g., for a local Bot API server).
- `WithNumWorkers(n int)`: Number of parallel workers for processing updates (default: 1).
- `WithClientTimeout(d time.Duration)`: HTTP client timeout.
- `WithHTTPClient(client *http.Client)`: Use a custom HTTP client (e.g., for proxy).
- `WithPanicRecovery()`: Enables middleware for recovering from panics in handlers.
- `WithTestEnv()`: Switches the bot to the Telegram test environment.

#### Long Polling

- `WithAllowedUpdates(ut ...types.UpdateType)`: List of update types you want to receive.
- `WithPollTimeout(d time.Duration)`: Long polling timeout (default: 1 minute).
- `WithPollLimit(limit int)`: Maximum number of updates retrieved per request (1-100).
- `WithPollOffset(offset int64)`: Offset for receiving updates.
- `WithUpdatesChanCap(cap int)`: Capacity of the internal update channel.

#### Webhook

- `WithWebhookURL(url string)`: Sets the webhook path based on the provided URL.
- `WithWebhookSecretToken(token string)`: Secret token for verifying requests from Telegram.

#### Debugging and Errors

- `WithDebugHandler(h HandlerFunc)`: Custom handler for logging every incoming update.
- `WithErrorHandler(h bot.ErrorHandler)`: Handler for system and network level errors.

## 🏗️ Framework Structure

### 💎 bot.Context — Heart of the Framework

`bot.Context` is passed to all handlers, filters, and middleware. It combines `context.Context` (Go), the bot object, and the current update.

- **`tg.Bot`**: Access to all Telegram API methods (`SendMessage`, `GetMe`, etc.).
- **`tg.Update`**: Data of the current event.
- **`tg.ChatID() / tg.SenderID()`**: Convenient helpers for getting Chat/User ID.
- **`tg.Set(key, val) / tg.Get(key)`**: Passing data between middleware and handlers.

### 🕹️ bot.Bot — Control

The main object for configuring and managing the bot's lifecycle.

- **Registration**: `b.Handler()`, `b.Group()`, `b.Use()`.
- **Startup**: `b.StartPolling()`, `b.StartWebhook()`.
- **API**: All Telegram Bot API methods in PascalCase.

### 📊 types.Update — Data

A complete typed Telegram update structure. Contains helper methods for quick content checking (`HasText()`, `IsCommand()`, `IsCallback()`), which automatically find the message inside the update.

## ⚡ Handlers

A Handler in Telegnom is a configurable unit consisting of:

1. **Logic** (`HandlerFunc`): What to do.
2. **Trigger** (`OnMessage`, `OnCallback`...): When to do it.
3. **Filters**: Clarifying conditions (e.g., only commands).
4. **Local Middleware**: Logic "around" the handler.

### Important Rules:

- **Order Matters**: The router checks handlers in the order they were registered. The first matching handler stops further search (First Match Wins).
- **Registration before Startup**: All handlers and groups must be fully configured **before** calling `StartPolling()` or `StartWebhook()`. Dynamic addition of handlers after startup is not supported.

```go
// 1. Register specific handlers first
b.Handler(handleStart).OnMessage(filters.Command("start"))

// 2. Then more general ones
b.Handler(echoHandler).OnMessage(filters.HasText)

// 3. (Optional) Handler with local middleware
b.Handler(secretHandler).OnMessage(filters.Command("admin")).Use(AuthMiddleware)

// 4. And only at the very end start the bot
b.StartPolling()
```

> **Tip**: If you use multiple triggers (e.g., `.OnMessage().OnEditedMessage()`), the handler will be registered in the corresponding processing queues.

See detailed examples (including Dependency Injection and multi-triggers) in [examples/handlers/README.md](examples/handlers/README.md).

### 🔍 Filters

Filters determine whether a handler should be executed. They work on an "all or nothing" principle: if you passed multiple filters, all of them must return `true`.

In Telegnom, the filtering system is built hierarchically (from high-level to low-level):

1. **`filters` Package**: Ready-made functions for registration (`f.Command("start")`, `f.Text("hello")`) and aliases to update methods (`f.HasPhoto`, `f.IsPrivate`).
2. **`types.Update` Methods**: Logic related to the update itself (`tg.Update.HasText()`, `tg.Update.IsCallback()`). Convenient as they find the necessary message inside the update themselves.
3. **`types.Message` Methods**: Checks for specific message fields (`msg.IsCommand()`, `msg.HasVideo()`).
4. **Custom Filters**: Any function with the signature `func(tg *bot.Context) bool`.

> Detailed guide and examples of all types of filters are in [examples/filters/README.md](examples/filters/README.md).

```go
// Example using different levels of filtering
b.Handler(h1).OnMessage(filters.Command("start")) // filters package
b.Handler(h2).OnMessage(func(tg *bot.Context) bool {
    return tg.Update.HasPhoto() // Update methods
})
```

## 📂 Handler Groups

Groups allow you to combine multiple handlers and apply common middleware to them (e.g., for authorization, logging, or access restriction).

```go
// Create a group
adminGroup := b.Group()

// Apply middleware only to this group
adminGroup.Use(AdminCheckMiddleware)

// All handlers in the group automatically inherit the middleware
adminGroup.Handler(adminStats).OnMessage(filters.Command("stats"))
adminGroup.Handler(adminBan).OnMessage(filters.Command("ban"))
```

> **Important**: Groups cannot be nested. Filters are applied only at the level of specific handlers within the group.

Details and an example of implementing admin panel protection can be found in [examples/groups/README.md](examples/groups/README.md).

## 🧅 Middleware

Middleware allows executing code BEFORE and AFTER event processing. Telegnom uses an "onion" architecture:

1. **Global** (`b.Use`): For all events.
2. **Group** (`group.Use`): For handlers in a group.
3. **Local** (`handler.Use`): For a specific single handler.

### How it works:

```go
func MyMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
    return func(tg *bot.Context) {
        // Code BEFORE handler (IN)
        next(tg)
        // Code AFTER handler (OUT)
    }
}
```

> **Stopping the chain**: If you do not call `next(tg)` and return, execution stops, and the handler will not be called. This is ideal for authorization.

Detailed examples and description of data passing via context can be found in [examples/middleware/README.md](examples/middleware/README.md).

## 🔄 Update Lifecycle

Understanding how Telegnom processes each update is critical for writing predictable code. The entire process can be visualized as a pipeline.

### Processing Scheme

```text
    [ Telegram Update ]
           ⬇
    [ Global Middleware (IN) ]  <-- (1. Global Layer)
           ⬇
       (Router looks for handler) ─┐
           ⬇                      │ (if filters don't match,
    [  Filters Check  ] ───────────┘  look for next handler)
           ⬇
       (Handler found!)           <-- (2. Decision Point)
           ⬇
    [ Group Middleware (IN) ]     <-- (3. Group Layer)
           ⬇
    [ Local Middleware (IN) ]     <-- (4. Local Layer)
           ⬇
    ┌───────────────────────┐
    │   HANDLER FUNCTION    │     <-- (5. Your Code)
    └───────────────────────┘
           ⬇
    [ Local Middleware (OUT) ]
           ⬇
    [ Group Middleware (OUT) ]
           ⬇
    [ Global Middleware (OUT) ]
```

### Nesting Visualization (Trace)

Since middleware works on the "matryoshka" (or onion) principle, the execution order looks like this. Imagine we have logging at each level:

```text
[Global]  Started processing update #100
  [Group]   Checked Admin rights... OK
    [Local]   Prepared user context
      >>> HANDLER: Executing logic for /start command
    [Local]   Cleanup context
  [Group]   Saved admin stats
[Global]  Finished processing (Took 2ms)
```

### Flow Control

You can stop processing at different stages, but the semantics will differ:

#### 1. Filters — "Face Control"
Filters are checked **before** launching group middleware and the handler itself.
- If a filter returns `false` -> The router simply looks for the **next** suitable handler.
- **Scenario**: "This command is only available in private chats". If written in a group — the bot simply ignores it (or finds another handler).

#### 2. Middleware — "Security Guard"
Middleware starts when a handler is already found.
- If middleware does not call `next(tg)` -> Processing **stops** completely. No other handlers are searched.
- **Scenario**: "Command found, but user has no rights". Middleware writes "Access denied" and returns.

## 🏁 Running the Bot

Telegnom supports two ways of receiving updates: **Long Polling** and **Webhook**.
[Telegram Documentation](https://core.telegram.org/bots/api#making-requests).

- **Long Polling**: The bot periodically polls Telegram servers for new events. This is the simplest way, not requiring an external IP address or SSL certificate. Ideal for local development and small bots.
- **Webhook**: Telegram sends an HTTP request to your server as soon as an event occurs. This method is more efficient and responds faster to messages but requires a public URL and configured HTTPS. Recommended for high-load projects.

### 📡 Long-Polling Mode

The `StartPolling()` method starts an infinite loop of polling the Telegram API for updates. It is blocking — program execution stops until interruption (e.g., Ctrl+C).

Options (optional) for Long Polling ([TG API DOC](https://core.telegram.org/bots/api#getupdates)):

- `WithPollLimit(limit int)`: Sets the maximum number of updates retrieved per request (from 1 to 100).
- `WithPollOffset(offset int64)`: Sets the initial update ID for polling. Positive value — start from a specific ID; negative (e.g., -1) — get only the last update.
- `WithPollTimeout(d time.Duration)`: Sets the timeout for a single polling request.
- `WithAllowedUpdates(ut ...types.UpdateType)`: Sets the types of updates the bot will receive.

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

	// 2. Initialize bot
	b, err := bot.NewBot(ctx, token, bot.WithPanicRecovery())
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	// 3. Register handlers
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Running in Long Polling mode!",
		})
	}).OnMessage()

	// 4. Start infinite polling loop
	log.Println("Bot started in Long Polling mode...")
	b.StartPolling()
}
```

> `bot.StartPolling()` will return an error if an outgoing webhook is configured. We recommend calling `bot.DeleteWebhook()` before starting.

### 🔗 Webhook Mode

In Webhook mode, Telegram sends updates to your server itself. This eliminates delays and saves resources. Telegnom offers three ways to work with webhooks:

1.  **Built-in HTTP Server**: `bot.StartWebhook(":8080")` — simplest way (usually used behind Nginx | Ngrok).
2.  **Built-in HTTPS Server**: `bot.StartWebhookTLS(":443", "cert.pem", "key.pem")` — if the bot faces the internet directly.
3.  **External Server**: `bot.WebhookHandler()` — for integration into an existing web application (Gin, Echo, etc.).

> **Important**: Before starting the bot in this mode, you must manually call `tg.Bot.SetWebhook()` (once when changing mode or parameters).

#### Usage Example (Built-in Server)

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

	// 1. Initialize bot with webhook configuration
	webhookURL := "https://my-domain.com/secret-path"
	secretToken := "my-secret-token"

	b, _ := bot.NewBot(ctx, token,
		bot.WithWebhookURL(webhookURL),
		bot.WithWebhookSecretToken(secretToken),
	)

	// 2. Register webhook in Telegram (done once or when settings change)
	_, err := b.SetWebhook(&bot.SetWebhookParams{
		URL:         webhookURL,
		SecretToken: secretToken,
	})
	if err != nil {
		log.Fatalf("Error registering webhook: %v", err)
	}

	// 3. Register handlers
	b.Handler(func(tg *bot.Context) {
		tg.Bot.SendMessage(&bot.SendMessageParams{
			ChatID: tg.ChatID(),
			Text:   "Received via Webhook!",
		})
	}).OnMessage()

	// 3. Start built-in server
	log.Println("Bot started in Webhook mode on port :8080...")
	if err := b.StartWebhook(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
```

## 📁 File Handling

For sending files (photos, videos, documents), the `types.InputFile` structure is used. You can send files by URL, FileID, disk path, or from `io.Reader`.

Detailed documentation: [docs/InputFile.md](docs/InputFile.md).

```go
// Sending photo from disk
_, _ = tg.Bot.SendPhoto(&bot.SendPhotoParams{
    ChatID: tg.ChatID(),
    Photo:  &types.InputFile{FilePath: "images/cat.jpg"},
    Caption: "Kitty",
})
```

## ⚠️ Error Handling

In Telegnom, error handling is divided into two levels: **system errors** and **Telegram API errors**. This is done so you can flexibly react to network problems separately from bot logic errors.

### Error Types

1. **System Errors**: Internet connection issues, DNS, JSON parsing errors. These errors are returned as standard Go errors.
2. **Telegram API Errors**: Errors returned by Telegram itself (e.g., "Chat not found", "Bot was blocked by the user", "Flood control exceeded"). They are returned as type `*bot.TGError`.

### How to Handle API Errors

In Go, use `errors.As` to check the error type.

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
			// This is an error from Telegram itself
			fmt.Printf("API Error: %d - %s\n", tgErr.Code, tgErr.Description)

			// Special convenience methods:
			if tgErr.IsFlood() {
				retryAfter := tgErr.GetRetryAfter()
				fmt.Printf("Need to wait %d seconds\n", retryAfter)
			}

			if newChatID, ok := tgErr.MigrateTo(); ok {
				fmt.Printf("Chat migrated to supergroup with ID: %d\n", newChatID)
			}
		} else {
			// This is a system error (network, proxy, etc.)
			fmt.Printf("System Error: %v\n", err)
		}
	}
}
```

### Centralized Handler (ErrorHandler)

By default, the library uses a built-in handler that logs system errors to the console via `log.Printf`. You can replace it with your own or disable it completely.

- **Default:** Enabled (built-in logger).
- **Custom Handler:** Pass a function to `bot.WithErrorHandler(myFunc)`.
- **Disable:** Pass `nil` to `bot.WithErrorHandler(nil)`.

Using a custom handler is useful for sending reports to Sentry, writing to files, or notifying an admin in Telegram.

```go
opts := []bot.Option{
    // Custom handler
    bot.WithErrorHandler(func (err error) {
        log.Printf("Critical error: %v", err)
    }),
}
// Or disable (logs will not be output)
// bot.WithErrorHandler(nil),
b, _ := bot.NewBot(ctx, token, opts...)
```

> **Important**: API errors (`TGError`) do **not** fall into `ErrorHandler` automatically when calling methods like `SendMessage`. They are returned to you directly for manual handling. `ErrorHandler` receives network-level errors and errors that occur inside the framework itself (e.g., when receiving updates).

## 📝 Logging and Debugging

Telegnom has built-in debugging tools to simplify development and monitoring.

### Viewing Incoming Updates (DebugHandler)

If your handler is not triggering, it is likely due to filters. To see exactly what data is coming from Telegram, use `WithDebugHandler`.

```go
opts := []bot.Option{
    // Use built-in logger (outputs Update in JSON)
    bot.WithDebugHandler(nil), // nil disables standard handler
    
    // Or write your own
    bot.WithDebugHandler(func(tg *bot.Context) {
        log.Printf("ID: %d, Text: %s", tg.Update.ID, tg.Update.Message.Text)
    }),
}
```

### Panic Recovery (PanicRecovery)

To prevent a single error in handler code from crashing the entire bot, always use `WithPanicRecovery()`. It intercepts panics, logs them via `ErrorHandler`, and allows the bot to continue working.

```go
b, _ := bot.NewBot(ctx, token, bot.WithPanicRecovery())
```

## 👤 User State Management

Coming Soon...

## 📂 Project Structure

Brief overview of Telegnom source code organization:

- **`bot/`**: 🧠 Framework core. Contains bot logic, context management, and routing.
    - **`methods.go`**: 🛠️ Implementation of all Telegram Bot API methods (`SendMessage`, `GetMe`, etc.).
    - **`methods_params.go`**: 📝 Parameter structures for API methods (e.g., `SendMessageParams`).
- **`types/`**: 📊 Data types. Full set of structures corresponding to Telegram Bot API objects (Update, Message, User, etc.).
- **`filters/`**: 🔍 Package with built-in filters for flexible handler trigger configuration.
- **`examples/`**: 🎓 Usage examples. Ready-made recipes for working with middleware, groups, webhooks, and more.
- **`docs/`**: 📖 Documentation. Additional materials and specifications for individual components.

## Handling Special Update Types

Coming Soon...

## Working with Keyboards

Coming Soon...

## Typical Usage Scenarios

Coming Soon...

## Testing

Coming Soon...
