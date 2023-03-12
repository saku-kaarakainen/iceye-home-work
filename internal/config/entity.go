package config

type Symbol struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Combination struct {
	Name         string `json:"name"`
	Value        int    `json:"value"`
	Regex_format string `json:"regex_format"`
}

type Config struct {
	Symbols      []Symbol
	Combinations []Combination
}
