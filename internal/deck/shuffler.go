package deck

import (
	"math/rand"
	"time"
)

func ShuffleDeck(deck *Deck) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle((len(deck.Cards)), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}
