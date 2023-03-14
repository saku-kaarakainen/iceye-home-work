package deck

import "fmt"

func (d Deck) String() string {
	hand := handToString(d)
	hname := d.Hand
	return fmt.Sprintf("[%s] (%s)", hand, hname)
}

func handToString(d Deck) string {
	var hand string
	for _, card := range d.Cards {
		hand += card.String()
	}

	return hand
}
