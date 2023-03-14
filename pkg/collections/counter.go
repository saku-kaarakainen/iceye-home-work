package collections

func CountSameLetters(value string) map[rune]int {
	m := make(map[rune]int)
	
	for _, c := range value {
		// Check if a key is present in the map
		if val, ok := m[c]; ok {
			m[c] = val + 1
		} else {
			m[c] = 1
		}
	}

	return m
}