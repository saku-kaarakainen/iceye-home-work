package collections

import t "larvis/pkg/types"

// key is value of the array.
// value is the count of such values
func Count[T comparable](array []T) []t.KeyValuePair[T, int] {
	dict := make([]t.KeyValuePair[T, int], 0)

	for _, dictKey := range array {
		// Check if a key is present in the map
		if ok, count := KeyExists(dict, dictKey); ok {
			Set(dict, dictKey, count+1)
		} else {
			dict = append(dict, t.KeyValuePair[T, int]{
				Key:   dictKey,
				Value: 1,
			})
		}
	}

	return dict
}
