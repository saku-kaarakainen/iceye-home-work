package game

import (
	"larvis/internal/actor"
	"larvis/internal/deck"
)

func (g *game) DealCards(
	cardsPerPlayer int,
	tableDeck *deck.Deck,
	players ...*actor.Actor) {

	for i := 0; i < cardsPerPlayer; i++ {
		for _, player := range players {

			// Remove the last element from the slice
			last := tableDeck.Cards[len(tableDeck.Cards)-1]       // save the last element
			tableDeck.Cards = tableDeck.Cards[:len(tableDeck.Cards)-1] // remove the last element

			player.Deck.Cards = append(player.Deck.Cards, last)
		}
	}
}
