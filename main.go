package main

import (
	"fmt"
	"larvis/internal/actor"
	"larvis/internal/card"
	"larvis/internal/deck"
	"larvis/internal/game"
)

func main() {
	cfg, _ := LoadConfig("./configs/config.json")
	// convert symbols into map, so they can be passed over different domains
	symbolsMap := card.MustCvrtSymbolsToMap(cfg.Domains.Card.Symbols)

	// initialize the game table
	g := game.NewGame(cfg.Domains.Game)
	g.AvailableCards = card.AssembleCards(cfg.Domains.Card)

	// invite actors, one dealer and two players
	g.Dealer = actor.CreateDealer()
	actor.GetAvailableCards(&g.Dealer, g.AvailableCards)

	g.Player1 = actor.CretePlayer1()
	g.Player2 = actor.CreateLarvis()

	// the game loop

	deck.ShuffleDeck(&g.Dealer.Deck)

	// Dealer is the Actor
	// deals the card from himself (source) to players (dests)
	actor.Deal(
		cfg.Domains.Game.CardsPerPlayer,
		&g.Dealer,
		&g.Player1,
		&g.Player2)

	deck.CalculatePoints(
		cfg.Domains.Deck,
		symbolsMap,
		&g.Player1.Deck)

	deck.CalculatePoints(
		cfg.Domains.Deck,
		symbolsMap,
		&g.Player2.Deck)

	actor.ShowCards(g.Player1)
	actor.ShowCards(g.Player2)

	winner := g.GetWinner()
	if (winner == nil) {
		fmt.Println("TIE")
	} else {
		fmt.Printf("Winner: %s\n", winner.Name)
	}

	actor.TakeCardsBack(
		&g.Dealer,
		&g.Player1,
		&g.Player2)

	fmt.Println("Thank you for playing poker with LARVIS")
}
