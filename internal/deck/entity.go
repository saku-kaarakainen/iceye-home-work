package deck

import "larvis/internal/card"

type Deck struct {
	Cards []card.Card
	Points [2]int
	Hand Hand
}
