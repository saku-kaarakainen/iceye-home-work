package deck

import (
	"larvis/internal/card"
)

type Combination struct {
	Name         string `json:"name"`
	Value        int    `json:"value"`
	Regex_format string `json:"regex_format"`
}

type DeckConfig struct {
	Combinations []Combination `json:"combinations"`
}

type Deck struct {
	Cards []card.Card
}
