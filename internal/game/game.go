package game

import (
	"fmt"
	"larvis/internal/config"
)

type Card struct {
}

type Player struct {
	cards []Card
}

type game struct {
	player1 Player
	larvis  Player

	Deck []Card
}

func NewGame(cfg config.RootConfig) *game {
	return &game{}
}

func assembleDeck(cfg config.RootConfig) {
	cfg.DeckConfigs.Colors
}

func (g *game) ShuffleDeck() {
	fmt.Println("shuffle deck")
}

func (g *game) DealCards() {
	fmt.Println("dealing cards")
	fmt.Println("You got: 2 of spades")
}

func (g *game) Play() bool {
	fmt.Println("play game")
	fmt.Println("Winner is: Mr. Larvis")

	// restart if true
	return true
}
