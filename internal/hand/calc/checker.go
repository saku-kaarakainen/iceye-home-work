package calc

import col "larvis/pkg/collections"

func GetChecker(hand HandType) func(map[int]int) bool {
	return map[HandType]func(map[int]int) bool{
		FourOfKind: IsFourOfKind,
		Triple:     IsTripple,
		FullHouse:  IsFullHouse,
		TwoPairs:   IsTwoPairs,
		Pair:       IsPair,
		HighCard:   IsHighCard,
	}[hand]
}

func IsFourOfKind(data map[int]int) bool {
	return col.HasValue(data, 4)
}

func IsFullHouse(data map[int]int) bool {
	return col.HasValue(data, 3) && col.HasValue(data, 2)
}

func IsTripple(data map[int]int) bool {
	return col.HasValue(data, 3)
}

func IsTwoPairs(data map[int]int) bool {
	return col.HasNKeysWithSameValue(data, 2, 2)
}

func IsPair(data map[int]int) bool {
	return col.HasValue(data, 2)
}

func IsHighCard(data map[int]int) bool {
	return true
}
