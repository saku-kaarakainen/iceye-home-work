package main

import (
	"fmt"
	"larvis/internal/card"
	"larvis/internal/game"
	"larvis/internal/hand"
)

func main() {
	cfg, _ := LoadConfig("./configs/config.json")
	symbolMap := card.MustCvrtSymbolsToMap(cfg.Domains.Card.Symbols)

	handCount := 2
	game := game.NewGame()
	for i := 0; i < handCount; i++ {
		game.Hands[i] = hand.New(
			cfg.Domains.Hand,
			symbolMap,
			fmt.Sprintf("Hand %d", i+1),
		)
		game.Hands[i].ReadHand()
		if game.Hands[i].IsValidHand() {
			game.Hands[i].CalculatePoints()
			break
		}
	}

	winner := game.GetWinner()
	fmt.Print("Winner: %s", winner.Name)
	fmt.Println()
}
