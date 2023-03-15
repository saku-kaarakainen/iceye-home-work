package game

import (
	"fmt"
	"larvis/internal/hand"
)

func (g *game) GetWinner() *hand.Hand {
	fmt.Println("Calculating winner... Hands", g.Hands)

	sp1 := g.Hands[0].Score[0]
	sp2 := g.Hands[1].Score[0]

	if sp1 == sp2 {
		fmt.Println("Both have same hands, calculating high card...")

		// Both have same hand
		sp1 = g.Hands[0].Score[1]
		sp2 = g.Hands[1].Score[1]
	}

	fmt.Println("PLAYER 1: ", sp1)
	fmt.Println("PLAYER 2: ", sp2)

	if sp1 == sp2 {
		return nil
	} else if sp1 > sp2 {
		return &g.Hands[0]
	} else {
		return &g.Hands[1]
	}
}
