package hand

import (
	col "larvis/pkg/collections"
	conv "larvis/pkg/convert"
	m "larvis/pkg/math"
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

func GetMethod(hand HandType) func(map[rune]int, string) (bool, int) {
	return map[HandType]func(map[rune]int, string) (bool, int){
		FourOfKind: IsFourOfKind,
		Triple:     IsTripple,
		FullHouse:  IsFullHouse,
		TwoPairs:   IsTwoPairs,
		Pair:       IsPair,
		HighCard:   IsHighCard,
	}[hand]
}

func IsFourOfKind(cardValues map[rune]int, value string) (bool, int) {
	data := col.CountSameLetters(value)
	res, run := col.HasValue(data, 4)
	return res, cardValues[run]
}

func IsFullHouse(cardValues map[rune]int, value string) (bool, int) {
	data := col.CountSameLetters(value)
	msk, key := col.HasValue(data, 3)
	lsk, _ := col.HasValue(data, 2)

	return msk && lsk, cardValues[key]
}

func IsTripple(cardValues map[rune]int, value string) (bool, int) {
	data := col.CountSameLetters(value)
	res, run := col.HasValue(data, 3)
	return res, cardValues[run]
}

func IsTwoPairs(cardValues map[rune]int, value string) (bool, int) {
	data := col.CountSameLetters(value)
	res, keys := col.HasNKeysWithSameValue(data, 2, 2)
	vals := conv.ConvertWithMapping(keys, cardValues)

	return res, m.FindMax(vals)
}

func IsPair(cardValues map[rune]int, value string) (bool, int) {
	data := col.CountSameLetters(value)
	res, run := col.HasValue(data, 2)
	return res, cardValues[run]
}

func IsHighCard(cardValues map[rune]int, value string) (bool, int) {
	keys := []rune(value)
	vals := conv.ConvertWithMapping(keys, cardValues)

	// This logic expects that the other hand types are checked first.
	// then there are nothing else left, so it must be high card.
	return true, m.FindMax(vals)
}