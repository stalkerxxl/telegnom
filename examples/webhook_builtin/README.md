# Built-in Webhook Example

This example demonstrates how to use the built-in webhook server (using `StartWebhook`), suitable for use with ngrok or a reverse proxy. It includes configuration for `WEBHOOK_URL` and security best practices like `SecretToken`.

## Usage

1. Set your `BOT_TOKEN` and `WEBHOOK_URL` environment variables.
2. Run the example: `go run main.go`
3. The bot will automatically call `SetWebhook` and start listening on `:8080`.

## Important Nuances

### 1. Graceful Shutdown
The library is designed to shut down gracefully. When the program receives a termination signal (like `Ctrl+C` or `SIGTERM`), the `context.Context` is canceled. This triggers the following sequence:
- The HTTP server stops accepting new connections.
- The bot stops accepting new updates from Telegram.
- The system waits for active workers to finish processing their current updates.
- `StartWebhook` returns, allowing your program to exit cleanly.

### 2. Webhook Path
By using `bot.WithWebhookURL("https://your-domain.com/my-bot-path")`, the library automatically configures the internal router to listen only on `/my-bot-path`. This is more secure than listening on the root `/`.

### 3. Secret Token
The `SecretToken` (via `WithWebhookSecretToken`) is highly recommended. It ensures that only Telegram can send requests to your endpoint. The library automatically validates the `X-Telegram-Bot-Api-Secret-Token` header for you.

### 4. Registration
`SetWebhook` is called in this example for convenience. In a production environment, you might only need to call it once when your URL or certificate changes, rather than on every startup.