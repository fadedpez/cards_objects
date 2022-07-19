package entities

import "fmt"

type Card struct {
	Suit  string
	Value string
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}
