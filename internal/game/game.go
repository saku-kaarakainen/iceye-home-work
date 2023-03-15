package game

import "larvis/internal/hand"

type Getter interface {
	GetWinner() *hand.Hand
}

type game struct {
	Hands [2]hand.Hand
}

func NewGame() *game {
	return &game{
		Hands: [2]hand.Hand{},
	}
}

func GetWinner(g Getter) *hand.Hand {
	return g.GetWinner()
}