package actor

import "larvis/internal/deck"

func CretePlayer1() Actor {
	return Actor{
		Name:  "Player 1",
		Role:  PLAYER,
		Deck: deck.Deck{},
	}
}

func CreateLarvis() Actor {
	return Actor{
		Name:  "Mr. Larvis",
		Role:  COMPUTER,
		Deck: deck.Deck{},
	}
}

func CreateDealer() Actor {
	return Actor{
		Name:  "Dealer",
		Role:  DEALER,
		Deck: deck.Deck{},
	}
}
