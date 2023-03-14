package deck

import (
	"sort"
	"strings"
)

func CalculatePoints(
	cfg DeckConfig,
	deck Deck,
) int {
	strval := deckToValueString(deck)
	return getMsPoint(cfg, strval)
}

func getMsPoint(
	cfg DeckConfig,
	strval string,
) int {
	// sort combinations by value, so that the most value combination
	// will be checked first
	sort.Sort(ByValueDesc(cfg.Combinations))
	for _, comb := range cfg.Combinations {
		isHand := GetMethod(Hand(comb.Method))

		if isHand(strval) {
			return comb.Value
		}
	}

	return 0
}

func deckToValueString(deck Deck) string {
	var symbols []string
	for _, card := range deck.Cards {
		symbols = append(symbols, card.Symbol.Name)
	}

	return strings.Join(symbols, "")
}
