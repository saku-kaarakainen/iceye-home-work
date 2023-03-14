package main

import (
	"fmt"
	"larvis/internal/actor"
	"larvis/internal/card"
	"larvis/internal/deck"
	"larvis/internal/game"
	"time"
)

func main() {
	cfg, _ := LoadConfig("./configs/config.json")

	// initialize the game table
	g := game.NewGame(cfg.Domains.Game)
	g.AvailableCards = card.AssembleCards(cfg.Domains.Card)

	// invite actors, one dealer and two players
	g.Dealer = actor.CreateDealer()
	actor.GetAvailableCards(&g.Dealer, g.AvailableCards)

	g.Player1 = actor.CretePlayer1()
	g.Player2 = actor.CreateLarvis()

	// the game loop
	for {
		deck.ShuffleDeck(&g.Dealer.Deck)

		// Dealer is the Actor
		// deals the card from himself (source) to players (dests)
		actor.Deal(
			cfg.Domains.Game.CardsPerPlayer,
			&g.Dealer,
			&g.Player1,
			&g.Player2)

		actor.ShowCards(g.Player1)
		pp := deck.CalculatePoints(cfg.Domains.Deck, g.Player1.Deck)

		actor.ShowCards(g.Player2)
		lp := deck.CalculatePoints(cfg.Domains.Deck, g.Player2.Deck)
	
		g.DeclareWinner()

		actor.TakeCardsBack(
			&g.Dealer,
			&g.Player1,
			&g.Player2)

		time.Sleep(5 * time.Second)
		// if !g.PlayAgain() {
		// 	break
		// }
	}

	fmt.Println("Thank you for playing poker with LARVIS")
}
