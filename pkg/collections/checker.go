package collections

// HasValue checks if the given map contains a value,
// and returns true and the key that has the value if found.
// Otherwise, it returns false and the zero value of T.
//
// Type T must be comparable, which means that
// the equality operator == must be defined for type T.
//
// Example:
//
//	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
//	found, key := HasValue(m, 2)
func HasValue[T comparable](m map[T]int, target int) bool {
	for _, v := range m {
		if v == target {
			return true
		}
	}

	return false
}

// HasNKeysWithSameValue checks if the given map has n keys
// that have the same value.
//
// The function returns true
// if there are exactly `count` keys with the same value,
// otherwise it returns false.
//
// Type T must be comparable, which means that
// the equality operator == must be defined for type T.
//
// Example:
//
//	m := map[rune]int{'a': 1, 'b': 2, 'c': 2}
//	result := HasNKeysWithSameValue(m, 2, 2)
//	// result is true because there are two keys ('b' and 'c')
//	// that have the value of 2
//
// Note:
//   - The function assumes that the map has no nil keys. If the map contains nil keys,
//     the behavior of this function is undefined.
//   - If the map has less than `count` keys with the same value, the function returns false.
func HasNKeysWithSameValue[T comparable](
	m map[T]int,
	value int,
	count int,
) bool {
	var c int
	for _, v := range m {
		if v == value {
			c++
		}
	}
	return c == count
}
