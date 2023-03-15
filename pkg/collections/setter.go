package collections

import t "larvis/pkg/types"


func Set[Key comparable, Value any](
	dict []t.KeyValuePair[Key, Value],
	key Key,
	value Value,
) {
	for i, kvp := range dict {
		if kvp.Key == key {
			dict[i].Value = value
			return
		}
	}
}