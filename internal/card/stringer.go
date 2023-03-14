package card

import "fmt"

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Color.Code, c.Symbol.Name)
}
