package game

import "larvis/internal/calc"

func (g *game) GetWinner() string {
	switch calc.FindBiggerScore(
		g.Hands[0].Score,
		g.Hands[1].Score,
	) {
	case calc.AisBigger:
		return "Hand 1"
	case calc.BisBigger:
		return "Hand 2"
	case calc.Tie:
	default:
		return "Tie"
	}

	panic("unreachable")
}
