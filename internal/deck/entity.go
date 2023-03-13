package deck

type Combination struct {
	Name         string `json:"name"`
	Value        int    `json:"value"`
	Regex_format string `json:"regex_format"`
}

type DeckConfig struct {
	Colors       []Color       `json:"colors"`
	Symbols      []Symbol      `json:"symbols"`
}


