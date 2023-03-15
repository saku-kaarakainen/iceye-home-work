package calc

import col "larvis/pkg/collections"

func (c *Calc) GetChecker(hand HandType) func() bool {
	return map[HandType]func() bool{
		FourOfKind: c.IsFourOfKind,
		Triple:     c.IsTripple,
		FullHouse:  c.IsFullHouse,
		TwoPairs:   c.IsTwoPairs,
		Pair:       c.IsPair,
		HighCard:   c.IsHighCard,
	}[hand]
}

func (c *Calc) IsFourOfKind() bool {
	return col.HasValue(c.data, 4)
}

func (c *Calc) IsFullHouse() bool {
	return col.HasValue(c.data, 3) && col.HasValue(c.data, 2)
}

func (c *Calc) IsTripple() bool {
	return col.HasValue(c.data, 3)
}

func (c *Calc) IsTwoPairs() bool {
	return col.HasNKeysWithSameValue(c.data, 2, 2)
}

func (c *Calc) IsPair() bool {
	return col.HasValue(c.data, 2)
}

func (c *Calc) IsHighCard() bool {
	return true
}
