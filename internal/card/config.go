package card

type SymbolConfig struct {
	Code  string `json:"code"`
	Val int    `json:"value"`
}

type CardConfig struct {
	Symbols []SymbolConfig `json:"symbols"`
}
