package calc

import (
	coll "larvis/pkg/collections"
	conv "larvis/pkg/convert"
	t "larvis/pkg/types"
	"sort"
)

// gets score and data
//
// data is passed as an array of key-value pairs,
// so that it's order can be determnistic.
//
// - score (1st item) - integer value of the hand. It's not necessary suitable to compare as-is, 
// because it does not take into account of the hand, e.g. if the hand is full house or triple.
//
// - data (2nd item)  - ordered data of the hand. first item ist the most valuable group. 
// Key is the value of the card and value is amount of such cards.
func GetScoreAndData(
	cfg map[rune]int,
	value string,
) (int, []t.KeyValuePair[int, int]) {
	data := toData(cfg, value)
	idx := toIndex(cfg, data)
	return idx, data
}

func toData(
	cfg map[rune]int,
	value string,
) []t.KeyValuePair[int, int] {
	// data as characters
	runes := []rune(value)

	// as their internal, configured values
	vals := conv.Map(runes, func(val rune) int {
		return cfg[val]
	})

	// sorted in descending order
	// eg. 27T9A -> AT972, and A333Q -> AQ333
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] > vals[j]
	})

	group := coll.Count(vals)

	sort.Slice(group, func(i, j int) bool {
		return group[i].Value > group[j].Value
	})

	return group
}

func toIndex(
	cfg map[rune]int,
	data []t.KeyValuePair[int, int],
) int {
	// flatten the array
	cards := flatten(data)
	i := sliceToInt(cards, 14)

	return i
}

func flatten(data []t.KeyValuePair[int, int]) []int {
	arr := make([]int, 0)
	for _, kvp := range data {
		for i := 0; i < kvp.Value; i++ {
			arr = append(arr, kvp.Key)
		}
	}
	return arr
}

func sliceToInt(s []int, base int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= base
	}
	return res
}
