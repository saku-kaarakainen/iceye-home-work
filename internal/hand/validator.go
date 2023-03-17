package hand

import "fmt"

func (h *Hand) IsValidHand() error {
	if len(h.Cards) != h.cfg.HandSize {
		return fmt.Errorf("hand size:%d, your hand:%d", h.cfg.HandSize, len(h.Cards))
	}

	firstChar := h.Cards[0]
	allSame := true

	for i, c := range h.Cards {
		_, exists := h.symCfg[c]
		if !exists {
			return fmt.Errorf("unknown symbol: %c", c)
		}

		if h.Cards[i] != firstChar {
			allSame = false
		}
	}

	if allSame {
		// can't accept, because typical poker deck does not allow that.
		return fmt.Errorf("all cards are the same")
	}

	return nil
}
