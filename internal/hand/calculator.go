package hand

import "sort"

func (h *Hand) CalculatePoints() {
	// sort combinations by value, so that the most value combination
	// will be checked first
	sort.Sort(ByValueDesc(h.cfg.Combinations))

	for _, comb := range h.cfg.Combinations {
		isHand := GetMethod(HandType(comb.Method))
		isIt, point := isHand(h.symCfg, h.Cards)

		if isIt {
			h.Score = [2]int{comb.Value, point}
			h.ScoreName = HandType(comb.Name)
			break
		}
	}
}
