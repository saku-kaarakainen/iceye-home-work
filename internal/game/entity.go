package game

type Combination struct {
	Name         string `json:"name"`
	Value        int    `json:"value"`
	Regex_format string `json:"regex_format"`
}

type GameConfig struct {
	Combinations   []Combination `json:"combinations"`
	CardsPerPlayer int           `json:"cardsPerPlayer"`
}
