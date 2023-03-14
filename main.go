package main

import (
	"fmt"
	"larvis/internal/card"
	"larvis/internal/deck"
	"larvis/internal/game"
)

func main() {

	cfg, _ := LoadConfig("./configs/config.json")

	// convert symbols into map, so they can be passed over different domains
	// cards are allowed to pass deck domain,
	// because they have m2m relationship with the deck
	symbolsMap := card.MustCvrtSymbolsToMap(cfg.Domains.Card.Symbols)

	// initialize the game table
	g := game.NewGame(cfg.Domains.Game)

	hand1 := g.GetHand(
		cfg.Domains.Deck,
		symbolsMap,
	)

	hand2 := g.GetHand(
		cfg.Domains.Deck,
		symbolsMap,
	)

	deck.CalculatePoints(
		hand1,
		symbolsMap,
		&g.Player1)

	deck.CalculatePoints(
		hand2,
		symbolsMap,
		&g.Player2)

	winner := g.GetWinner()
	if winner == nil {
		fmt.Println("TIE")
	} else {
		fmt.Printf("Winner: %s\n", winner.Name)
	}

	fmt.Println("Thank you for playing poker with LARVIS")
}
