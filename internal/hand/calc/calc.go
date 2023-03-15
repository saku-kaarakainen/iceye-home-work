package calc

import (
	col "larvis/pkg/collections"
	conv "larvis/pkg/convert"
	"sort"
)

type Calc struct {
	cfg  map[rune]int
	data map[int]int
	hand HandType
}

func NewCalc(cfg map[rune]int) *Calc {
	return &Calc{cfg: cfg}
}

func (c *Calc) CalcData(value string) {
	// data as characters
	runes := []rune(value)

	// as their internal, configured values
	vals := conv.Map(runes, func(val rune) int {
		return c.cfg[val]
	})

	// sorted in descending order
	// eg. 27T9A -> AT972, and A333Q -> AQ333
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] > vals[j]
	})

	group := col.Count(vals)

	// set the data
	c.data = col.SortMapByValues(group)
}
