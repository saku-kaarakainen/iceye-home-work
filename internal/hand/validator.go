package hand

import "fmt"

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