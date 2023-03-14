package card

func AssembleCards(cfg CardConfig) []Card {
	cards := make([]Card, 0)

	for _, color := range cfg.Colors {
		for _, symbol := range cfg.Symbols {
			cards = append(cards, Card{Color: color, Symbol: symbol})
		}
	}

	return cards
}
