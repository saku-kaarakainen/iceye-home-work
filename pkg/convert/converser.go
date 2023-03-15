package convert

import "fmt"

// MapKeys transforms a map by applying a function to each key and returning
// a new map with the transformed keys and the original values.
//
// The function `f` takes a key of type `SrcK` from the source map and returns
// a key of type `DestK` for the destination map. The function is applied to
// every key in the source map.
//
// The types `SrcK`, `DestK`, and `Val` must be comparable. If the source map
// has a key type that is not comparable, this function will not compile.
//
// The returned map has the same value type as the input map.
//
// Example usage:
//
//	inputMap := map[string]int{"one": 1, "two": 2, "three": 3}
//	outputMap := MapKeys(inputMap, func(k string) string {
//	    return strings.ToUpper(k)
//	})
//	// Output: map[ONE:1 TWO:2 THREE:3]
func MapKeys[
	SrcK comparable,
	DestK comparable,
	Val any,
](
	src map[SrcK]Val,
	f func(SrcK) DestK,
) map[DestK]Val {
	newMap := make(map[DestK]Val)
	for k, v := range src {
		i := f(k)
		newMap[i] = v
	}
	return newMap
}

func Map[T, U any](ts []T, f func(T) U) []U {
    us := make([]U, len(ts))
    for i := range ts {
        us[i] = f(ts[i])
    }
    return us
}

// StringToRune takes a string as input and returns its first rune,
// as well as an error if the length of the string is not 1.
//
// Example:
//
//	r, err := StringToRune("a")
//	// r is 'a' (rune type)
//	// err is nil
//
// Note:
//   - The function assumes that the input string contains only one character.
//     If the input string is empty or contains more than one character,
//     the function returns an error.
func StringToRune(str string) (rune, error) {
	var err error
	str_len := len(str)
	if str_len != 1 {
		err = fmt.Errorf("string '%s' len is not 1, but %d", str, str_len)
	}

	return []rune(str)[0], err
}
