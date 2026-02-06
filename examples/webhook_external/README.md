# External Webhook Example (Echo)

This example demonstrates integration with an external web framework (`Echo`). It uses `WebhookHandler` to adapt the bot's update processing to the external server's routing and shows how to manage the lifecycle (graceful shutdown) of both the bot and the web server.

## Usage

1. Set your `BOT_TOKEN` and `WEBHOOK_URL` environment variables.
2. Ensure you have the Echo framework installed: `go get github.com/labstack/echo/v4`
3. Run the example: `go run main.go`

## Important Nuances

### 1. Proper Shutdown Sequence
When using an external server, you must manage the shutdown sequence carefully to avoid losing updates:
1. **Stop the Web Server**: Call `e.Shutdown()`. This ensures no new requests (updates) are accepted from Telegram.
2. **Cancel Bot Context**: This signals the internal workers to stop.
3. **Wait for Workers**: The `b.Run()` method (which should be running in a goroutine) will exit only after all buffered updates have been processed by the workers.

### 2. Integration via WebhookHandler
`b.WebhookHandler()` returns a standard `http.HandlerFunc`. You can wrap it for any framework:
- **Echo**: `e.POST("/path", echo.WrapHandler(b.WebhookHandler()))`
- **Gin**: `r.POST("/path", gin.WrapH(b.WebhookHandler()))`
- **Standard Library**: `http.HandleFunc("/path", b.WebhookHandler())`

### 3. Security
Even with an external server, you should use `bot.WithWebhookSecretToken()`. The `WebhookHandler` will still perform the verification of the secret token header automatically, even when called from within another framework.

### 4. Public URL and HTTPS
Telegram requires your webhook URL to be public and use HTTPS. If you are developing locally, use tools like **ngrok** to create a secure tunnel to your local port.