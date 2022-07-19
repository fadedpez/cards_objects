package entities

import "fmt"

type Card struct {
	Suit  string
	Value string
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}

func CreateDeck() []*Card {
	suits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	deck := make([]*Card, 0)

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, &Card{Suit: suit, Value: value})
		}
	}

	return deck

}
