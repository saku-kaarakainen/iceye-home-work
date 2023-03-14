package card


type Symbol struct {
	Code  string `json:"code"`
	Value int    `json:"value"`
}

type CardConfig struct {
	Symbols []Symbol `json:"symbols"`
}
