package types

// Dice represents an animated emoji that displays a random value. The bot API
// currently supports the following emoji: 1-6 for “🎲”, “🎯” and “🎳” base
// emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
//
// See https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}
