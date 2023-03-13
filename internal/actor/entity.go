package actor

import "larvis/internal/card"

type ActorConfig struct {
}

type Role string

const (
	PLAYER Role   = "player"
	COMPUTER Role = "computer"
	DEALER Role   = "dealer"
)

// Actor in the game
type Actor struct {
	Name  string
	Role Role
	Cards []card.Card
}
