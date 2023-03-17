package hand

import (
	"larvis/internal/calc"
	"larvis/pkg/types"
	"sort"
)

func (h *Hand) Calculate() {
	score, data := calc.GetScoreAndData(h.symCfg, h.Cards)
	h.Score = score
	h.Combination = h.getCombination(data)
}

func (h *Hand) getCombination(data []types.KeyValuePair[int, int]) Combination {
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
