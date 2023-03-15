package collections

import "sort"

// TODO: Make generic
func SortMapByValues(m map[int]int) map[int]int {
	// Create a slice to store the keys
	keys := make([]int, 0, len(m))
	vals := make([]int, 0, len(m))
	res := make(map[int]int)

	// Populate the keys slice
	for k, v := range m {
		keys = append(keys, k)
		vals = append(vals, v)
	}

	// Sort the keys slice
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	sort.Slice(vals, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	// populate the result
	for k, val := range vals {
		key := keys[k]
		res[key] = val
	}
	
	return res
}