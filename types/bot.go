package types

// BotName represents the bot's name.
// See https://core.telegram.org/bots/api#botname
type BotName struct {
	Name string `json:"name"`
}

// BotDescription represents the bot's description.
// See https://core.telegram.org/bots/api#botdescription
type BotDescription struct {
	Description string `json:"description"`
}

// BotShortDescription represents the bot's short description.
// See https://core.telegram.org/bots/api#botshortdescription
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}
