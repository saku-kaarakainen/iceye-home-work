package calc

type ScoreResult int

const (
	AisBigger ScoreResult = iota
	BisBigger
	Tie
)

func FindBiggerScore(
	a, b map[int]int,
	aCombVal, bCombVal int,
) ScoreResult {

	handRes := findBiggerByHandVal(aCombVal, bCombVal)
	if handRes != Tie {
		return handRes
	}

	return findBiggerByMapping(a, b)
}

func findBiggerByHandVal(aCombVal, bCombVal int) ScoreResult {
	if aCombVal > bCombVal {
		return AisBigger
	}

	if bCombVal > aCombVal {
		return BisBigger
	}

	return Tie
}

func findBiggerByMapping(a, b map[int]int) ScoreResult {
	if len(a) != len(b) {
		panic("mismatch in map lengths")
	}

	aKeys := make([]int, 0)
	for aKey, _ := range a {
		aKeys = append(aKeys, aKey)
	}

	bKeys := make([]int, 0)
	for bKey, _ := range b {
		bKeys = append(bKeys, bKey)
	}

	for i := 0; i < len(aKeys); i++ {
		aCard := aKeys[i]
		bCard := bKeys[i]

		if aCard > bCard {
			return AisBigger
		}
		if aCard < bCard {
			return BisBigger
		}
	}

	return Tie
}
