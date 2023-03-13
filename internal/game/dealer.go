package game

import (
	"larvis/internal/actor"
	"larvis/internal/deck"
)

func (g *game) DealCards(
	cardsPerPlayer int,
	deck *deck.Deck,
	players ...*actor.Actor) {

	for i := 0; i < cardsPerPlayer; i++ {
		for _, player := range players {

			// Remove the last element from the slice
			last := deck.Cards[len(deck.Cards)-1]       // save the last element
			deck.Cards = deck.Cards[:len(deck.Cards)-1] // remove the last element

			player.Cards = append(player.Cards, last)
		}
	}
}
