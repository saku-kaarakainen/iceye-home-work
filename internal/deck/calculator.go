package deck

import (
	"fmt"
	"larvis/internal/card"
	"regexp"
	"sort"
	"strings"
)

func CalculatePoints(
	cfg DeckConfig,
	deck Deck,
) int {
	str := deckToValueString(deck.Cards)
	p := findHighestCombination(cfg.Combinations, str)
	fmt.Printf("Calculating points for deck: %s, got %d points.\n",
		str, p)

	return 0
}

func findHighestCombination(
	combination []Combination,
	value string,
) int {
	sort.Sort(ByValueDesc(combination))
	for _, comb := range combination {
		regex := regexp.QuoteMeta(comb.Regex_format)

		fmt.Printf("Testing value '%s'. with %s\n", value, comb.Regex_format)
		re := regexp.MustCompile(regex)

		if re.MatchString(value) {
			fmt.Printf(" !!! Match with %d points\n", comb.Value)
			return comb.Value
		} else {
			fmt.Printf(" - No match for %d points\n", comb.Value)
		}
	}

	return 0
}

func deckToValueString(cards []card.Card) string {
	var symbols []string
	for _, card := range cards {
		symbols = append(symbols, card.Symbol.Name)
	}

	return strings.Join(symbols, "")
}
