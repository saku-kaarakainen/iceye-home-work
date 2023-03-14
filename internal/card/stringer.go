package card

import "fmt"

func (c Card) String() string {
	// if you need to know the color as well
	// return fmt.Sprint(c.Color.Code) + c.Symbol.Code)
	return fmt.Sprint(c.Symbol.Code)
}
