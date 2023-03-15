package score


func GetIndex(cards string) {

}


func (h *Hand) getMethod(hand HandType) func(string) bool {
	return map[HandType]func(string) bool {
		FourOfKind: h.IsFourOfKind,
		Triple:     h.IsTripple,
		FullHouse:  h.IsFullHouse,
		TwoPairs:   h.IsTwoPairs,
		Pair:       h.IsPair,
		HighCard:   h.IsHighCard,
	}[hand]
}

// sort string and convert them into card's internal value
func (h *Hand) sortAndVal(value string) []int {
	// data as characters
	runes := []rune(value)

	// as their internal, configured values
	vals := conv.Map(runes, func(val rune) int {
		return h.symCfg[val]
	})

	// sorted in descending order
	// eg. 27T9A -> AT972
	sort.Slice(vals, func(i, j int) bool { 
		return vals[i] > vals[j] 
	})


	return vals
}