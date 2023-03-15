package calc

type ScoreResult int

const (
	AisBigger ScoreResult = iota
	BisBigger
	Tie
)

func FindBiggerScore(a, b map[int]int) ScoreResult {
	return Tie
}
