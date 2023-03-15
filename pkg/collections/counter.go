package collections

import t "larvis/pkg/types"

// key is value of the array.
// value is the count of such values
func Count[T comparable](array []T) []t.KeyValuePair[T, int] {
	tArr := make([]t.KeyValuePair[T, int], 0)

	for i, val := range array {
		// Check if a key is present in the map
		if KeyExists(tArr, val) {
			tArr[i].Value = tArr[i].Value + 1
		} else {
			tArr[i].Value = 1
		}
	}

	return tArr
}

