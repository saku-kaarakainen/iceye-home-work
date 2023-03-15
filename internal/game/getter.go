package game

import "larvis/internal/hand"

func (g *game) GetWinner() *hand.Hand {
	sp1 := g.Hands[0].Score[0]
	sp2 := g.Hands[1].Score[0]

	if sp1 == sp2 {
		// Both have same hand
		sp1 = g.Hands[0].Score[1]
		sp2 = g.Hands[1].Score[1]
	}

	if sp1 == sp2 {
		return nil
	} else if sp1 > sp2 {
		return &g.Hands[0]
	} else {
		return &g.Hands[1]
	}
}
