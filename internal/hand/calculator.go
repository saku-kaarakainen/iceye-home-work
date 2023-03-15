package hand

import (
	"fmt"
	"larvis/internal/hand/calc"
	"sort"
)

func (h *Hand) CalculatePoints() {
	h.calc.CalcData(h.Cards)
	comb := h.getCombination()

	// set base points, for example
	// high card 0
	// pair 100
	// full house 400
	// 1000 is just arbitrary big number
	score := comb.Value * 100000000000
	fmt.Println("score", score)
}

func (h *Hand) getCombination() Combination {
	// sort combinations by value, so that the most value combination
	// will be checked first
	sort.Sort(ByValueDesc(h.cfg.Combinations))
	for _, comb := range h.cfg.Combinations {
		hand := calc.HandType(comb.Method)
		isHand := h.calc.GetChecker(hand)()
		if isHand {
			return comb
		}
	}

	panic("no combination found")
}
