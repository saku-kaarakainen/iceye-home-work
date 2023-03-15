package hand

import (
	"larvis/internal/hand/calc"
)

type Reader interface {
	ReadHand()
}

type Validator interface {
	IsValidHand() error
}

type Calculator interface {
	CalculatePoints()
}

type Hand struct {
	Name  string
	Cards string
	Score [2]int

	cfg    HandConfig
	symCfg map[rune]int
	calc   *calc.Calc
}

func New(
	cfg HandConfig,
	symCfg map[rune]int,
	name string,
) Hand {
	return Hand{
		Name:   name,
		cfg:    cfg,
		symCfg: symCfg,
		calc:   calc.NewCalc(symCfg),
	}
}

func ReadHand(hand Reader) error {
	hand.ReadHand()
	return nil
}

func ValidateHand(hand Validator) error {
	return hand.IsValidHand()
}

func CalculatePoints(hand Calculator) {
	hand.CalculatePoints()
}
