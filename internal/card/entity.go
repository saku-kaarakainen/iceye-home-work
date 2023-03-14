package card

type Color struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Symbol struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type CardConfig struct {
	Colors  []Color  `json:"colors"`
	Symbols []Symbol `json:"symbols"`
}

type Card struct {
	Color  Color
	Symbol Symbol
}