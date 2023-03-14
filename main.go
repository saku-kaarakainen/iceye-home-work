package main

import (
	"fmt"
	"larvis/internal/card"
	"larvis/internal/game"
	"larvis/internal/hand"
)

func main() {
	cfg, _ := LoadConfig("./configs/config.json")
	fmt.Println(cfg)
	symbolMap := card.MustCvrtSymbolsToMap(cfg.Domains.Card.Symbols)
	fmt.Println(symbolMap)
	handCount := 2
	game := game.NewGame()
	for i := 0; i < handCount; i++ {
		for {
			game.Hands[i] = hand.New(
				cfg.Domains.Hand,
				symbolMap,
				fmt.Sprintf("Hand %d", i+1),
			)
			game.Hands[i].ReadHand()

			err := game.Hands[i].IsValidHand() 
			if err == nil {
				fmt.Println(err)
				continue
			}

			game.Hands[i].CalculatePoints()
			break
		}
	}

	winner := game.GetWinner()
	if winner == nil {
		fmt.Print("Winner: TIE")
	} else {
		fmt.Print("Winner: %s", winner.Name)
	}
	fmt.Println()
}
