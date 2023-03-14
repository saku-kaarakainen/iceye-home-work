package card

import "larvis/pkg/convert"

func MustCvrtSymbolsToMap(symbols []Symbol) map[rune]int {
	symbolMap := make(map[rune]int)

	for _, symbol := range symbols {
		run, err := convert.StringToRune(symbol.Code)
		if err != nil {
			panic(err)
		}

		symbolMap[run] = symbol.Value
	}
	return symbolMap
}

