package game

import (
	"fmt"
	"larvis/internal/actor"
	"larvis/internal/deck"
)

type game struct {
	Player1 actor.Actor
	Player2 actor.Actor
	Dealer  actor.Actor

	Deck deck.Deck
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

func (g *game) TakeCarsBackToDeck(
	deck *deck.Deck,
	players ...*actor.Actor) {
	for _, player := range players {
		deck.Cards = append(deck.Cards, player.Deck.Cards...)
		player.Deck.Cards = nil
	}
}
