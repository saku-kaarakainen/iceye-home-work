package actor

import "fmt"

func ShowCards(actor Actor) {
	fmt.Printf("%s cards: %v (hand: %s)\n",
		actor.Name,
		actor.Deck.Cards,
		actor.Deck.Hand,
	)
}
