package calc

type HandType string

const (
	FourOfKind HandType = "FOUROFAKIND"
	Triple     HandType = "TRIPLE"
	FullHouse  HandType = "FULLHOUSE"
	TwoPairs   HandType = "TWOPAIRS"
	Pair       HandType = "APAIR"
	HighCard   HandType = "HIGHCARD"
)
