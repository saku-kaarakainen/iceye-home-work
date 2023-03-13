package card

import "fmt"

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

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Color.Code, c.Symbol.Name)
}
