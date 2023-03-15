package hand

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Hand struct {
	Name      string
	Cards     string
	Score     [2]int
	ScoreName HandType

	cfg    HandConfig
	symCfg map[rune]int
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
	}
}

func (h *Hand) ReadHand() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Give hand for %s:", h.Name)
	scanner.Scan()

	h.Cards = scanner.Text()
}

func (h *Hand) IsValidHand() error {
	if len(h.Cards) != h.cfg.HandSize {
		return fmt.Errorf("hand size:%d, your hand:%d", h.cfg.HandSize, len(h.Cards))
	}

	for _, c := range h.Cards {
		_, exists := h.symCfg[c]
		if !exists {
			return fmt.Errorf("unknown symbol: %c", c)
		}
	}

	return nil
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
			break
		}
	}
}
