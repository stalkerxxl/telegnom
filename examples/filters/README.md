# Filters System in Telegnom

Telegnom provides a flexible and hierarchical filtering system to help you route updates to the correct handlers.
A Filter is simply a function with the signature: `func(tg *bot.Context) bool`.

If a filter returns `true`, the handler chain continues. If `false`, the handler is skipped.

## The 3 Layers of Filters

### 1. The `filters` Package (High-Level)
This is what you will use 90% of the time. It contains:
- **Factories**: Functions that return a filter based on arguments (e.g., `filters.Command("start")`, `filters.Text("hello")`).
- **Aliases**: Shortcuts to `Update` methods (e.g., `filters.HasPhoto`, `filters.IsPrivate`).

```go
import f "github.com/stalkerxxl/telegnom/filters"

// Only runs for /start command
b.Handler(h).OnMessage(f.Command("start")) 
// Only runs for photos in private chats
b.Handler(h).OnMessage(f.HasPhoto, f.IsPrivate) 
```

### 2. `types.Update` Methods (Mid-Level)
The `types.Update` struct has many helper methods. They are "smart" because they automatically look for the `EffectiveMessage`. This means `HasText()` works for `Message`, `EditedMessage`, and `ChannelPost` without you checking each field.

You often use these inside custom filters or within handlers.

```go
func MyFilter(tg *bot.Context) bool {
    // Works regardless of whether it's an edited message or a new one
    return tg.Update.HasText() && tg.Update.IsPrivate() 
}
```

### 3. `types.Message` Methods (Low-Level)
The `types.Message` struct has methods that check the specific fields of that message instance. These are useful when you already have a `*types.Message` variable and want to check its properties.

```go
func handler(tg *bot.Context) {
    msg := tg.Update.EffectiveMessage()
    if msg.IsCommand() {
        // ...
    }
    if msg.HasSticker() {
        // ...
    }
}
```

## Creating Custom Filters

Creating a custom filter is as simple as writing a function.

```go
func AdminOnly(tg *bot.Context) bool {
    adminID := int64(123456789)
    return tg.SenderID() == adminID
}

// Usage
b.Handler(secretHandler).OnMessage(AdminOnly)
```

## Logic (AND / OR)

- **AND**: Implicit. When you pass multiple filters to `OnMessage(f1, f2)`, **ALL** of them must return true.
- **OR**: You need to implement this manually in a custom filter.

```go
// Example of OR logic: Command /start OR /help
func StartOrHelp(tg *bot.Context) bool {
    return f.Command("start")(tg) || f.Command("help")(tg)
}
```
