package hand

import (
	"larvis/internal/hand/calc"
	"sort"
)

func (h *Hand) Calculate() {
	h.Score = calc.GetData(h.symCfg, h.Cards)
	h.Combination = h.getCombination(h.Score)
}

func (h *Hand) getCombination(data map[int]int) Combination {
	// sort combinations by value, so that the most value combination
	// will be checked first
	sort.Sort(ByValueDesc(h.cfg.Combinations))
	for _, comb := range h.cfg.Combinations {
		hand := calc.HandType(comb.Method)
		isHand := calc.GetChecker(hand)(data)
		if isHand {
			return comb
		}
	}

	// happens only if there are issues in config values.
	panic("no combination found")
}
