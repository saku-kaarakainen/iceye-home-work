package deck

type Combination struct {
	Name   string `json:"name"`
	Value  int    `json:"value"`
	Method string `json:"method"`
}

type DeckConfig struct {
	Combinations []Combination `json:"combinations"`
}
