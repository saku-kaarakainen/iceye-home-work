package hand

import (
	"bufio"
	"fmt"
	"larvis/internal/card"
	"os"
	"sort"
)

type Hand struct {
	Name  string
	Cards string
	Score [2]int
	ScoreName HandType

	cfg HandConfig
	symCfg map[rune]int
}

func New(
	cfg HandConfig,
	symCfg map[rune]int,
	name string,
	) Hand {
	return Hand{
		Name: name, 
		cfg: cfg,
		symCfg: symCfg,
	}
}

func (h *Hand) ReadHand() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Give hand: ")
	scanner.Scan()

	h.Cards = scanner.Text()
}

func (h *Hand) IsValidHand() bool {
	if len(h.Cards) != h.cfg.HandSize {
		return false
	}

	for _, c := range h.Cards {
		_, exists := h.symCfg[c]
		if !exists {
			return false
		}
	}

	return true
}

func (h *Hand) CalculatePoints() {
	// sort combinations by value, so that the most value combination
	// will be checked first
	sort.Sort(ByValueDesc(h.cfg.Combinations))

	for _, comb := range h.cfg.Combinations {
		isHand := GetMethod(HandType(comb.Method))
		isIt, point := isHand(h.symCfg, h.Cards)

		if isIt {
			h.Score = [2]int{comb.Value, point}
			h.ScoreName = HandType(comb.Name)
		}
	}
}
