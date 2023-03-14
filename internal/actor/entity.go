package actor

// - Game can refer Actor (or Dealer)
// - Actor, Non-Dealer can refer Hand 
//     - Dealer can refer Deck
// - Deck/Hand can refer Card
import "larvis/internal/deck"

type Role string

const (
	PLAYER   Role = "player"
	COMPUTER Role = "computer"
	DEALER   Role = "dealer"
)

// Actor in the game
type Actor struct {
	Name   string
	Role   Role
	Deck   deck.Deck
	Points int
}
