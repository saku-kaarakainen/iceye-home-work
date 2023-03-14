package deck

import c "larvis/pkg/collections"

// hand could be separate domain.
type Hand string

const (
	FourOfKind Hand = "FOUROFAKIND"
	Triple     Hand = "TRIPLE"
	FullHouse  Hand = "FULLHOUSE"
	TwoPairs   Hand = "TWOPAIRS"
	Pair       Hand = "APAIR"
	HighCard   Hand = "HIGHCARD"
)

func GetMethod(hand Hand) func(string) bool {
	return map[Hand]func(string) bool{
		FourOfKind: IsFourOfKind,
		Triple:     IsTripple,
		FullHouse:  IsFullHouse,
		TwoPairs:   IsTwoPairs,
		Pair:       IsPair,
		HighCard:   IsHighCard,
	}[hand]
}

func IsFourOfKind(value string) bool {
	return c.HasValue(c.CountSameLetters(value), 4)
}

func IsFullHouse(value string) bool {
	res := c.CountSameLetters(value)
	return c.HasValue(res, 3) && c.HasValue(res, 2)
}

func IsTripple(value string) bool {
	return c.HasValue(c.CountSameLetters(value), 3)
}

func IsTwoPairs(value string) bool {
	return c.HasNKeysWithSameValue(c.CountSameLetters(value), 2, 2)
}

func IsPair(value string) bool {
	return c.HasValue(c.CountSameLetters(value), 2)
}

func IsHighCard(value string) bool {
	// High card is trivial because
	// it is always true if there is no other hand.
	return true
}
