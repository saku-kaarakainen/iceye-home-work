package hand

import (
	"bufio"
	"fmt"
	"os"
)

func (h *Hand) ReadHand() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Give hand for %s:", h.Name)
	scanner.Scan()

	h.Cards = scanner.Text()
}