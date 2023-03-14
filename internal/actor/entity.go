package actor

import "larvis/internal/deck"

type ActorConfig struct {
}

type Role string

const (
	PLAYER   Role = "player"
	COMPUTER Role = "computer"
	DEALER   Role = "dealer"
)

// Actor in the game
type Actor struct {
	Name string
	Role Role
	Deck deck.Deck
}
