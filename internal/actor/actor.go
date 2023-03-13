package actor

import "larvis/internal/card"

func CretePlayer1() Actor {
	return Actor{
		Name:  "Player 1",
		Role:  PLAYER,
		Cards: []card.Card{},
	}	
}

func CreateLarvis() Actor {
	return Actor{
		Name:  "Mr. Larvis",
		Role:  COMPUTER,
		Cards: []card.Card{},
	}
}

func CreateDealer() Actor {
	return Actor{
		Name:  "Dealer",
		Role:  DEALER,
		Cards: []card.Card{},
	}
}