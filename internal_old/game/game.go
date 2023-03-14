package game

import (
	"bufio"
	"fmt"
	"larvis/internal/card"
	"larvis/internal/hand"
	"os"
)

// TODO: Should have struct?
// Maybe should be deleted
type game struct {
	Player1 hand.Deck
	Player2 hand.Deck

	AvailableCards []card.Card
	config         GameConfig
}

// Initializes the game
func NewGame(config GameConfig) *game {
	return &game{
		config: config,
	}
}

// TODO: Correct domain?
func (g *game) GetHand(
	cfg deck.DeckConfig,
	symbolsMap map[rune]int,
) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Give hand: ")
		scanner.Scan()

		text := scanner.Text()
		if deck.IsValidHand(cfg, symbolsMap, text) {
			return text
		} else {
			fmt.Printf("I'm sorry, but '%s' is not proper hand for the game.", text)
			fmt.Println()
		}
	}
}

func (g *game) GetWinner() *deck.Deck {
	// significant points
	sp1 := g.Player1.Points[0]
	sp2 := g.Player2.Points[0]

	if sp1 == sp2 {
		sp1 = g.Player1.Points[1]
		sp2 = g.Player2.Points[1]
	}

	if sp1 == sp2 {
		return nil
	} else if sp1 > sp2 {
		return &g.Player1
	} else {
		return &g.Player2
	}
}
