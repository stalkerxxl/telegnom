# Handlers in Telegnom

A Handler in Telegnom is more than just a function. It is a configuration unit that consists of:
1.  **Logic**: The `HandlerFunc` that executes your code.
2.  **Triggers**: The update types (e.g., `OnMessage`, `OnCallbackQuery`) that activate it.
3.  **Conditions**: The filters (e.g., `filters.Command`) that must pass.
4.  **Local Middleware**: Specific behaviors applied only to this handler.

## Lifecycle of a Handler

When an update arrives, Telegnom checks handlers in the order they were registered.
For a handler to be executed, it must pass through this funnel:

```
[ Update Arrives ]
       |
       v
[ Trigger Check ] (Is it a Message? Callback?) -> NO: Skip
       |
       v
[ Filters Check ] (Is it a Command? Private chat?) -> NO: Skip
       |
       v
[ Local Middlewares (IN) ] -> (e.g. Logging, RateLimit)
       |
       v
[ Handler Function ] -> YOUR CODE HERE
       |
       v
[ Local Middlewares (OUT) ]
```

## Features

### Multi-Triggers
You can register the same handler function for multiple events.

```go
// React to both new and edited messages
b.Handler(myFunc).OnMessage().OnEditedMessage()
```

### Local Middleware
Add middleware that runs *only* for this specific handler.

```go
b.Handler(adminAction).OnMessage(f.Command("ban")).Use(AuthMiddleware)
```

### Dependency Injection (Closures)
Avoid global variables by using closures to inject dependencies like databases or services.

```go
func MyHandler(db *Database) bot.HandlerFunc {
    return func(tg *bot.Context) {
        user := db.FindUser(tg.SenderID())
        // ...
    }
}
```

## Organizing Many Handlers (Struct Pattern)

For larger projects with many handlers, it is recommended to group related handlers into structs. This allows you to:
- **Group by topic**: e.g., `UserHandler`, `AdminHandler`, `PaymentHandler`.
- **Inject dependencies**: Pass services, database connections, or configurations during struct initialization.
- **Keep code clean**: Each struct can live in its own file or package.

```go
type UserHandler struct {
    service *UserService
}

func (h *UserHandler) GetProfile(tg *bot.Context) {
    // ... use h.service here
}

// In main.go:
userHandler := &UserHandler{service: userService}
b.Handler(userHandler.GetProfile).OnMessage(f.Command("profile"))
```

See [examples/handlers/struct_handlers.go](struct_handlers.go) for a full implementation of this pattern.

## Best Practices

1.  **Order Matters**: Register specific handlers (like commands) *before* general handlers (like "catch-all" text handlers).
2.  **Keep it clean**: If your handler logic is long, move the function to a separate file or package.
3.  **Use Groups**: If you have many handlers with the same middleware (e.g., Auth), use `b.Group()` instead of adding `.Use()` to each one.
