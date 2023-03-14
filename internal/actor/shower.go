package actor

import "fmt"

func ShowCards(actor Actor) {
	fmt.Printf("%s: %v\n", actor.Name, actor.Deck)
}
