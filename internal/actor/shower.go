package actor

import "fmt"

func ShowCards( actor Actor) {
	fmt.Printf("%s cards: %v\n", actor.Name, actor.Cards)
}