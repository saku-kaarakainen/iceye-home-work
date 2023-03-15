package hand

import "sort"

func (h *Hand) CalculatePoints() {
	var score int
	i, hand := h.getHandIndex()

	// add base points
	length := len(h.cfg.Combinations) // 6
	score = i * length // for example: (full house) 5 * 6 = 30

	// add additional points
	score += 
}

func (h *Hand) getHandIndex() (int, HandType) {
	// sort combinations by value, so that the most value combination
	// will be checked first
	var hand HandType
	sort.Sort(ByValueDesc(h.cfg.Combinations))
	for i, comb := range h.cfg.Combinations {
		hand = HandType(comb.Method)
		getIsHand := h.getMethod(hand)
		
		if getIsHand(h.Cards) {
			return len(h.cfg.Combinations) - 1 - i, hand
		}
	}

	return 0, hand
}
