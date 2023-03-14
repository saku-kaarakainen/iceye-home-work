package card

type Color struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Symbol struct {
	Code  string `json:"code"`
	Value int    `json:"value"`
}

type CardConfig struct {
	Colors  []Color  `json:"colors"`
	Symbols []Symbol `json:"symbols"`
}
