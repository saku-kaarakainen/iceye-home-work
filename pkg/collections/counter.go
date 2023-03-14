package collections

// CountSameLetters returns a map where the keys are runes
// and the values are the number of occurrences of each rune
// in the input string.
//
// Example:
//
//	result := CountSameLetters("hello")
//	// result is a map with the following key-value pairs:
//	// {'h': 1, 'e': 1, 'l': 2, 'o': 1}
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
