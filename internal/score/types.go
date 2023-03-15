package score

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

func getMethod(hand HandType) func(string) (bool, int) {
	return map[HandType]func(string) (bool, int){
		FourOfKind: IsFourOfKind,
		Triple:     IsTripple,
		FullHouse:  IsFullHouse,
		TwoPairs:   IsTwoPairs,
		Pair:       IsPair,
		HighCard:   IsHighCard,
	}[hand]
}

func IsFourOfKind(value string) bool {
	data := col.CountSameLetters(value)
	return col.HasValue(data, 4)
}

func IsFullHouse(value string) (bool, int) {
	cards := h.sortAndVal(value)
	data := col.Count(cards)

	msk, key := col.HasValue(data, 3)

	lsk, _ := col.HasValue(data, 2)

	return msk && lsk, key
}

func IsTripple(value string) (bool, int) {
	data := h.calcData(value)
	res, key := col.HasValue(data, 3)
	return res, key
}

func IsTwoPairs(value string) (bool, int) {
	data := h.calcData(value)
	res, keys := col.HasNKeysWithSameValue(data, 2, 2)

	return res, m.FindMax(keys)
}

func IsPair(value string) (bool, int) {
	data := h.calcData(value)
	res, key := col.HasValue(data, 2)
	return res, key
}

func IsHighCard(value string) (bool, int) {
	data := h.calcData(value)
	keys := []rune(value)
	vals := conv.ConvertWithMapping(keys, cardValues)

	// This logic expects that the other hand types are checked first.
	// then there are nothing else left, so it must be high card.
	return true, m.FindMax(vals)
}
