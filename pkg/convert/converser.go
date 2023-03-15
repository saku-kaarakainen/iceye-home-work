package convert

import "fmt"

// Map applies a mapping function to each element of the input slice and returns the resulting slice of mapped values.
//
// The type parameters T and U define the types of the input and output elements respectively.
//
// The mapping function f takes an element of type T as its input and returns an element of type U.
//
// Example:
//   input := []int{1, 2, 3}
//   output := Map(input, func(x int) int { return x * 2 })
//   // output is now []int{2, 4, 6}
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
