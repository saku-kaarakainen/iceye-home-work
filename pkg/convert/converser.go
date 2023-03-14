package convert

import "fmt"

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

func StringToRune(str string) (rune, error) {
	var err error
	str_len := len(str)
	if str_len != 1 {
		err = fmt.Errorf("string '%s' len is not 1, but %d", str, str_len)
	}

	return []rune(str)[0], err
}
