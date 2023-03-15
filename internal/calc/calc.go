package calc

import (
	col "larvis/pkg/collections"
	conv "larvis/pkg/convert"
	"sort"
)

func GetData(
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

	group := col.Count(vals)

	return col.SortMapByValues(group)
}
