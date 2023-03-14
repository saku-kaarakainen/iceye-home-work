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
	var hands [2]hand.Hand

	for i := 0; i < handCount; i++ {
		hands[i] = hand.New(
			cfg.Domains.Hand,
			symbolMap,
			fmt.Sprintf("Hand %d", i+1),
		)
		hands[i].ReadHand()
		if hands[i].IsValidHand() {
			hands[i].CalculatePoints()
			break
		}
	}

	winner := game.GetWinner()
	fmt.Print("Winner: %s", winner.Name)
	fmt.Println()
}
