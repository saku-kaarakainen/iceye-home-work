package game

import (
	"fmt"
	"larvis/internal/actor"
	"larvis/internal/card"
)

type game struct {
	player1 actor.Actor
	larvis  actor.Actor
	dealer  actor.Actor

	Deck []card.Card
}

func NewGame() *game {

	return &game{}
}

func (g *game) Play() bool {
	fmt.Println("play game")
	fmt.Println("Winner is: Mr. Larvis")

	// restart if true
	return true
}
