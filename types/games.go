package types

// Game represents a game. Use BotFather to create and edit games, their short
// names will act as unique identifiers.
//
// See https://core.telegram.org/bots/api#game
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

// GameHighScore represents one row of the high scores table for a game.
//
// See https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}

// CallbackGame represents a placeholder, currently holds no information.
//
// See https://core.telegram.org/bots/api#callbackgame
type CallbackGame struct{}
