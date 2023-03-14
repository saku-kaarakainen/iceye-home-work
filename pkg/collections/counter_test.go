package collections

import "testing"

// TestStrAnalyzer is a unit test for the CountSameLetters function.
func TestStrAnalyzer(t *testing.T) {
	testCases := []struct {
		input string
		want  map[rune]int
	}{
		{
			"larvis", map[rune]int{
				'l': 1,
				'a': 1,
				'r': 1,
				'v': 1,
				'i': 1,
				's': 1,
			},
		},
		{
			"banana", map[rune]int{
				'b': 1,
				'a': 3,
				'n': 2,
			},
		},
		{
			"apple", map[rune]int{
				'a': 1,
				'p': 2,
				'l': 1,
				'e': 1,
			},
		},
		{
			"!!!!", map[rune]int{
				'!': 4,
			},
		},
	}

	for _, tc := range testCases {
		got := CountSameLetters(tc.input)
		if !mapEquals(got, tc.want) {
			t.Errorf("CountSameLetters(%s) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func mapEquals(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}
