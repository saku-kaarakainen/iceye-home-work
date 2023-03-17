package hand

type Combination struct {
	Name   string `json:"name"`
	Value  int    `json:"value"`
	Method string `json:"method"`
}

type HandConfig struct {
	HandSize     int           `json:"handSize"`
	Combinations []Combination `json:"combinations"`
}
