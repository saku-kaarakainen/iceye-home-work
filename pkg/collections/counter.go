package collections

func Count[T comparable](array []T) map[T]int {
	m := make(map[T]int)

	for _, val := range array {
		// Check if a key is present in the map
		if c, ok := m[val]; ok {
			m[val] = c + 1
		} else {
			m[val] = 1
		}
	}

	return m
}

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
