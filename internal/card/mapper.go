package card

import "larvis/pkg/convert"

func MustCvrtSymbolsToMap(symbols []SymbolConfig) map[rune]int {
	symbolMap := make(map[rune]int)

	for _, symbol := range symbols {
		run, err := convert.StringToRune(symbol.Code)
		if err != nil {
			panic(err)
		}

		symbolMap[run] = symbol.Val
	}
	return symbolMap
}
