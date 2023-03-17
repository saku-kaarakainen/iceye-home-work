package collections

import t "larvis/pkg/types"

func KeyExists[Key comparable, Value any](
	dict []t.KeyValuePair[Key, Value],
	key Key,
) (bool, Value) {
	for _, kvp := range dict {
		if kvp.Key == key {
			return true, kvp.Value
		}
	}

	var empty Value
	return false, empty
}

// ContainsValue checks if the given key-value-pair array contains a value,
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
func ContainsValue[K comparable, V comparable](
	dict []t.KeyValuePair[K, V], 
	value V,
	) bool {
	for _, kvp := range dict {
		if kvp.Value == value {
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
func CountValues[K comparable, V comparable](
	dict []t.KeyValuePair[K, V], 
	val V,
) int {
	c := 0
	for _, v := range dict {
		if v.Value == val {
			c++
		}
	}
	return c
}
