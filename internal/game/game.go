package game

func GetWinner() *deck.Deck {
	// significant points
	sp1 := g.Player1.Points[0]
	sp2 := g.Player2.Points[0]

	if sp1 == sp2 {
		sp1 = g.Player1.Points[1]
		sp2 = g.Player2.Points[1]
	}

	if sp1 == sp2 {
		return nil
	} else if sp1 > sp2 {
		return &g.Player1
	} else {
		return &g.Player2
	}
}