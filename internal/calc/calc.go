package calc

import (
	coll "larvis/pkg/collections"
	conv "larvis/pkg/convert"
	"sort"
)

func GetScoreAndData(
	cfg map[rune]int,
	value string,
) (int, map[int]int) {
	data := toData(cfg, value)
	idx := toIndex(cfg, data)
	return idx, data
}

func toData(
	cfg map[rune]int,
	value string,
) map[int]int {
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

	return coll.SortMapByValues(group)
}

func toIndex(
	cfg map[rune]int,
	data map[int]int,
) int {
	// flatten the array
	cards := flatten(data)
	i := sliceToInt(cards, 14)

	return i
}

func flatten(data map[int]int) []int {
	arr := make([]int, 0)
	for key, value := range data {
		for i := 0; i < value; i++ {
			arr = append(arr, key)
		}
	}
	return arr
}

// 14 max per card
// i=0
// 		-> 14-0 -> 14
// 		-> 13

func sliceToInt(s []int, base int) int {
    res := 0
    op := 1
    for i := len(s) - 1; i >= 0; i-- {
        res += s[i] * op
        op *= base
    }
    return res
}
