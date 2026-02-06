# Middleware in Telegnom

Middleware allows you to execute code **before** and/or **after** a handler. It follows the "onion" architecture: each middleware wraps the next one, and the handler is at the center.

## The 3 Layers of Middleware

Telegnom executes middleware in this specific order:

1.  **Global Middleware**: Applied via `b.Use()`. Runs for every single update.
2.  **Group Middleware**: Applied via `group.Use()`. Runs if the matched handler belongs to the group.
3.  **Local Middleware**: Applied via `handler.Use()`. Runs only for that specific handler.

## Anatomy of a Middleware

A middleware is a function that takes a `HandlerFunc` and returns a `HandlerFunc`.

```go
func MyMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
    return func(tg *bot.Context) {
        // 1. LOGIC BEFORE HANDLER (IN)
        log.Println("Before")

        next(tg) // 2. CALL NEXT (Middleware or Handler)

        // 3. LOGIC AFTER HANDLER (OUT)
        log.Println("After")
    }
}
```

### Stopping the Chain
To interrupt the execution (e.g., if authentication fails), simply **do not call `next(tg)`** and return from the function.

```go
func AuthMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
    return func(tg *bot.Context) {
        if !isAuthorized(tg) {
            return // Chain stops here. Handler will NEVER be called.
        }
        next(tg)
    }
}
```

## Passing Data via Context
Since `bot.Context` embeds `context.Context`, you can use it to pass data down the chain safely.

```go
func UserMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
    return func(tg *bot.Context) {
        user := db.GetUser(tg.SenderID())
        tg.Set("user", user) // Save to context
        next(tg)
    }
}

func Handler(tg *bot.Context) {
    user := tg.Get("user").(*User) // Retrieve in handler
}
```

See [examples/middleware/main.go](main.go) for a full working example with all 3 layers.
