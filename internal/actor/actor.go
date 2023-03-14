package actor

import "larvis/internal/deck"

func CretePlayer1() Actor {
	return Actor{
		Name: "Hand 1",
		Role: PLAYER,
		Deck: deck.Deck{},
	}
}

func CreateLarvis() Actor {
	return Actor{
		Name: "Hand 2",
		Role: PLAYER,
		Deck: deck.Deck{},
	}
}

func CreateDealer() Actor {
	return Actor{
		Name: "Dealer",
		Role: DEALER,
		Deck: deck.Deck{},
	}
}
