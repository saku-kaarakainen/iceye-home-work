package game

type Hand string

const (
	FourOfKind Hand = "FOUROFAKIND"
	Triple     Hand = "TRIPLE"
	FullHouse  Hand = "FULLHOUSE"
	TwoPairs   Hand = "TWOPAIRS"
	Pair       Hand = "APAIR"
	HighCard   Hand = "HIGHCARD"
)

func IsFourOfKind(value string) bool {

}
