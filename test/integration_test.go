// integration tests

package main

import (
	"fmt"
	"larvis/internal/card"
	"larvis/internal/game"
	"larvis/internal/hand"
	"larvis/pkg/json"
	"testing"
)

// TestScoringToWinner tests that scoring is translated to winner correctly.
func TestScoringToWinner(t *testing.T) {
	// Test - load config
	cfg, _ := json.Load[RootConfig]("./../configs/config.json")
	symbolMap := card.MustCvrtSymbolsToMap(cfg.Domains.Card.Symbols)

	for _, test := range geTestCases() {
		testName := fmt.Sprintf("Test |%s|%s|%s|", test.Hand1, test.Hand2, test.Winner) + fmt.Sprintln()

		t.Run(testName, func(t *testing.T) {
			g := game.NewGame()
			// Test - init hands
			g.Hands[0] = hand.New(cfg.Domains.Hand, symbolMap, "Hand 1")
			g.Hands[1] = hand.New(cfg.Domains.Hand, symbolMap, "Hand 2")

			g.Hands[0].Cards = test.Hand1
			g.Hands[1].Cards = test.Hand2

			// Test - calculate points
			g.Hands[0].Calculate()
			g.Hands[1].Calculate()

			// Test - get winner
			haveWinner := g.GetWinner()

			if haveWinner != test.Winner {
				t.Errorf("winner - want:'%v', have:'%v'", test.Winner, haveWinner)
			}
		})
	}
}
