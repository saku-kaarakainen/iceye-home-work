package deck

func IsValidHand(
	cfg DeckConfig,
	symbolsMap map[rune]int,
	input string,
) bool {
	if len(input) != cfg.HandSize {
		return false
	}

	for _, c := range input {
		_, exists := symbolsMap[c]
		if !exists {
			return false
		}
	}

	return true
}