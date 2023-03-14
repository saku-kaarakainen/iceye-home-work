package game

import (
	"fmt"
	"larvis/internal/actor"
	"larvis/internal/card"
)

type game struct {
	Player1 actor.Actor
	Player2 actor.Actor
	Dealer  actor.Actor

	AvailableCards []card.Card
}

// Initializes the game
func NewGame(cfg GameConfig) *game {
	return &game{}
}

func (g *game) GetWinner() *actor.Actor {
	// significant points
	sp1 := g.Player1.Deck.Points[0]
	sp2 := g.Player2.Deck.Points[0]
	
	if sp1 == sp2 {
		sp1 = g.Player1.Deck.Points[1]
		sp2 = g.Player2.Deck.Points[1]
	}

	if sp1 == sp2 {
		return nil
	} else if sp1 > sp2 {
		return &g.Player1
	} else {
		return &g.Player2
	}
}

func (g *game) PlayAgain() bool {
	fmt.Println("play again")
	return false
}