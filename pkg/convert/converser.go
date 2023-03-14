package convert

import "fmt"

// ConvertWithMapping takes a slice of source type 'Src' and a map
// of type [Src]Dest as input, and returns a slice of type 'Dest'.
// Each element in the output slice is mapped from its corresponding
// element in the input slice using the given mapping.
//
// Type Src and Dest must be comparable, which means that
// the equality operator == must be defined for these types.
//
// Example:
//
//	src := []string{"a", "b", "c"}
//	mapping := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := ConvertWithMapping(src, mapping)
//	// result is []int{1, 2, 3}
//
// Note:
//   - The function assumes that each element in the input slice exists as a key
//     in the given mapping. If not, the function panics.
func ConvertWithMapping[Src comparable, Dest any](
	src []Src,
	mapping map[Src]Dest,
) []Dest {
	dest := make([]Dest, len(src))
	for i, r := range src {
		dest[i] = mapping[r]
	}
	return dest
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
