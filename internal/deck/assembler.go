package deck

import (
	"larvis/internal/card"
)

func AssembleDeck(cfg card.CardConfig) Deck {
	deck := Deck{}

	for _, color := range cfg.Colors {
		for _, symbol := range cfg.Symbols {
			deck.Cards = append(deck.Cards, card.Card{Color: color, Symbol: symbol})
		}
	}

	return deck
}