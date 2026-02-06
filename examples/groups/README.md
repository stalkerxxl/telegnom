# Handler Groups in Telegnom

Handler Groups allow you to group related handlers and apply shared middleware to them. This is the best way to implement features like:
- **Authentication**: Only allow certain users to access a set of commands.
- **Logging**: Log activity for a specific section of the bot.
- **Feature Toggling**: Enable or disable a group of commands at once.

## Key Features

1.  **Shared Middleware**: Use `group.Use(middleware)` to apply logic to all handlers in that group.
2.  **Order of Execution**: When an update arrives and a handler from a group is matched:
    1.  **Filters** are checked first.
    2.  **Group Middlewares** run next.
    3.  **Local Middlewares** run after that.
    4.  The **Handler Function** runs last.
3.  **First Match Wins**: Telegnom searches for handlers in the order they were registered, regardless of which group they belong to.

## Example: Admin Protection

```go
admin := b.Group()
admin.Use(AdminAuthMiddleware)

// Both of these will be protected by AdminAuthMiddleware
admin.Handler(handleStats).OnMessage(f.Command("stats"))
admin.Handler(handleBan).OnMessage(f.Command("ban"))
```

## Important Restrictions

- **No Nesting**: Groups cannot be nested (you cannot create a group inside another group).
- **No Group Filters**: Filters are applied only to individual handlers via `.OnMessage(filters...)`. If you need a "group filter", implement it as a middleware that calls `return` without calling `next(tg)`.

See [examples/groups/main.go](main.go) for a full working example.
