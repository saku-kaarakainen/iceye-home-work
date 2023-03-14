package deck

import (
	"sort"
	"strings"
)

func CalculatePoints(
	cfg DeckConfig,
	cardValues map[rune]int,
	deck *Deck,
) {
	copy := *deck
	strval := deckToValueString(copy)

	p, h := getMsPoint(cfg, cardValues, strval)
	deck.Points = p
	deck.Hand = h
}

func getMsPoint(
	cfg DeckConfig,
	cardValues map[rune]int,
	strval string,
) ([2]int, Hand) {
	// sort combinations by value, so that the most value combination
	// will be checked first
	sort.Sort(ByValueDesc(cfg.Combinations))

	for _, comb := range cfg.Combinations {
		isHand := GetMethod(Hand(comb.Method))
		isIt, point := isHand(cardValues, strval)

		if isIt {
			return [2]int{comb.Value, point}, Hand(comb.Name)
		}
	}

	return [2]int{0, 0}, ""
}

func deckToValueString(deck Deck) string {
	var symbols []string
	for _, card := range deck.Cards {
		symbols = append(symbols, card.Symbol.Code)
	}

	return strings.Join(symbols, "")
}
