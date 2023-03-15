package calc

import col "larvis/pkg/collections"
import t "larvis/pkg/types"

func GetChecker(hand HandType) func([]t.KeyValuePair[int, int]) bool {
	return map[HandType]func([]t.KeyValuePair[int, int]) bool{
		FourOfKind: IsFourOfKind,
		Triple:     IsTripple,
		FullHouse:  IsFullHouse,
		TwoPairs:   IsTwoPairs,
		Pair:       IsPair,
		HighCard:   IsHighCard,
	}[hand]
}

func IsFourOfKind(data []t.KeyValuePair[int, int]) bool {
	return col.ContainsValue(data, 4)
}

func IsFullHouse(data []t.KeyValuePair[int, int]) bool {
	return col.ContainsValue(data, 3) && col.ContainsValue(data, 2)
}

func IsTripple(data []t.KeyValuePair[int, int]) bool {
	return col.ContainsValue(data, 3)
}

func IsTwoPairs(data []t.KeyValuePair[int, int]) bool {
	return col.CountValues(data, 2) == 2
}

func IsPair(data []t.KeyValuePair[int, int]) bool {
	return col.ContainsValue(data, 2)
}

func IsHighCard(data []t.KeyValuePair[int, int]) bool {
	return true
}
