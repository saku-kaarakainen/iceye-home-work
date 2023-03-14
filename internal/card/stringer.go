package card

import "fmt"

func (c Card) String() string {
	return fmt.Sprint(string(c.Color.Code) + string(c.Symbol.Code))
}
