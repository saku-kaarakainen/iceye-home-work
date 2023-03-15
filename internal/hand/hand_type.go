package hand

import (
	col "larvis/pkg/collections"
	conv "larvis/pkg/convert"
	m "larvis/pkg/math"
	"sort"
)

type HandType string

const (
	FourOfKind HandType = "FOUROFAKIND"
	Triple     HandType = "TRIPLE"
	FullHouse  HandType = "FULLHOUSE"
	TwoPairs   HandType = "TWOPAIRS"
	Pair       HandType = "APAIR"
	HighCard   HandType = "HIGHCARD"
)

func (h *Hand) getMethod(hand HandType) func(string) (bool, int) {
	return map[HandType]func(string) (bool, int){
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


func (h *Hand) rank(cards []int) {
	// for example 	AAKKK
	// 				AA432
	// 				AKKK3

	// each card has reserved 5-bits, 
	// because there are 5 combinations
	// most valuable combination has highest index
	
	index := getIndex(cards)


}

func (h *Hand) IsFourOfKind(value string) bool {
	data := col.Count(h.sortAndVal(value))
	return col.HasValue(data, 4)
}

func (h *Hand) IsFullHouse(value string) (bool, int) {
	cards := h.sortAndVal(value)
	data := col.Count(cards)

	msk := col.HasValue(data, 3) && col.HasValue(data, 2)
}

func (h *Hand) IsTripple(value string) (bool, int) {
	data := h.calcData(value)
	res, key := col.HasValue(data, 3)
	return res, key
}

func (h *Hand) IsTwoPairs(value string) (bool, int) {
	data := h.calcData(value)
	res, keys := col.HasNKeysWithSameValue(data, 2, 2)

	return res, m.FindMax(keys)
}

func (h *Hand) IsPair(value string) (bool, int) {
	data := h.calcData(value)
	res, key := col.HasValue(data, 2)
	return res, key
}

func (h *Hand) IsHighCard(value string) (bool, int) {
	data := h.calcData(value)
	keys := []rune(value)
	vals := conv.ConvertWithMapping(keys, cardValues)

	// This logic expects that the other hand types are checked first.
	// then there are nothing else left, so it must be high card.
	return true, m.FindMax(vals)
}
