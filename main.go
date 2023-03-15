package main

import (
	"fmt"
	"larvis/internal/card"
	"larvis/internal/game"
	"larvis/internal/hand"
	"larvis/pkg/json"
)

func main() {
	cfg, _ := json.Load[RootConfig]("./configs/config.json")
	symbolMap := card.MustCvrtSymbolsToMap(cfg.Domains.Card.Symbols)

	game := game.NewGame()
	for i := 0; i < cfg.Domains.Game.PlayerCount; i++ {
		game.Hands[i] = hand.New(
			cfg.Domains.Hand,
			symbolMap,
			fmt.Sprintf("Hand %d", i+1),
		)

		for {
			game.Hands[i].ReadHand()

			err := game.Hands[i].IsValidHand()
			if err != nil {
				fmt.Println("I'm sorry I can't calculate the score for your hand, because ", err)
				continue
			}

			game.Hands[i].Calculate()
			break
		}
	}

	winner := game.GetWinner()
	if winner == nil {
		fmt.Print("Winner: TIE")
	} else {
		fmt.Printf("Winner: %s", winner.Name)
	}
	fmt.Println()
	fmt.Println("Thank you for playing poker with LARVIS!")
}
