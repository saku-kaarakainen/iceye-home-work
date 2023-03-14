package actor

import "larvis/internal/deck"

type Role string

const (
	PLAYER Role = "player"
	DEALER Role = "dealer"
)

// Actor in the game
type Actor struct {
	Name   string
	Role   Role
	Deck   deck.Deck
	Points int
}
