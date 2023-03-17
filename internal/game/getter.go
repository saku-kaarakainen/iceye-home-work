package game

func (g *game) GetWinner() string {
	// check hands
	if g.Hands[0].Combination.Value > g.Hands[1].Combination.Value {
		return g.Hands[0].Name
	}
	if g.Hands[0].Combination.Value < g.Hands[1].Combination.Value {
		return g.Hands[1].Name
	}

	// both have same hand, check score
	if g.Hands[0].Score > g.Hands[1].Score {
		return g.Hands[0].Name
	}
	if g.Hands[0].Score < g.Hands[1].Score {
		return g.Hands[1].Name
	}

	return "Tie"
}
