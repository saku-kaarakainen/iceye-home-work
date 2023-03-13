package card

type Color struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Symbol struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
