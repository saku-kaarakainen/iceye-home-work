package main

import (
	"fmt"
	"larvis/internal/actor"
	"larvis/internal/deck"
	"larvis/internal/game"
)

func main() {
	cfg, _ := LoadConfig("./configs/config.json")

	// initialize the game table
	g := game.NewGame(cfg.Domains.Game)
	g.Deck = deck.AssembleDeck(cfg.Domains.Card)

	// invite actors, one dealer and two players
	g.Dealer = actor.CreateDealer()
	g.Player1 = actor.CretePlayer1()
	g.Player2 = actor.CreateLarvis()

	// the game loop
	for {
		deck.ShuffleDeck(&g.Deck)
		g.DealCards(
			cfg.Domains.Game.CardsPerPlayer,
			&g.Deck,
			&g.Player1, &g.Player2)

		actor.ShowCards(g.Player1)
		deck.CalculatePoints(cfg.Domains.Deck, g.Player1.Deck)

		actor.ShowCards(g.Player2)
		deck.CalculatePoints(cfg.Domains.Deck, g.Player2.Deck)


		g.TakeCarsBackToDeck(
			&g.Deck,
			&g.Player1,
			&g.Player2)

		if !g.PlayAgain() {
			break
		}
	}

	fmt.Println("Thank you for playing poker with LARVIS")
}
