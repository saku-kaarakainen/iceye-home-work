package actor

import "larvis/internal/card"

func GetAvailableCards(
	actor *Actor,
	availableCards []card.Card) {
	actor.Deck.Cards = make([]card.Card, len(availableCards))
	copy(actor.Deck.Cards, availableCards)
	availableCards = nil
}

func Deal(
	cardsPerPlayer int,
	source *Actor,
	dests ...*Actor) {

	for i := 0; i < cardsPerPlayer; i++ {
		for _, dest := range dests {

			// Remove the last element from the slice
			last := source.Deck.Cards[len(source.Deck.Cards)-1]              // save the last element
			source.Deck.Cards = source.Deck.Cards[:len(source.Deck.Cards)-1] // remove the last element

			dest.Deck.Cards = append(dest.Deck.Cards, last)
		}
	}
}

func TakeCardsBack(
	dest *Actor,
	sources ...*Actor) {
	for _, src := range sources {
		dest.Deck.Cards = append(dest.Deck.Cards, src.Deck.Cards...)
		src.Deck.Cards = nil
	}
}
