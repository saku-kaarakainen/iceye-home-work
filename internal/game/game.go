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

func (g *game) DeclareWinner() {
	fmt.Println("declare winner")
}

func (g *game) PlayAgain() bool {
	fmt.Println("play again")
	return false
}