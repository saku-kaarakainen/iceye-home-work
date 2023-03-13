package game

import (
	"fmt"
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

func NewGame() *game {
	return &game{}
}

func assembleDeck() {

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
