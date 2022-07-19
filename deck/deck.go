package deck

import (
	"cards_objects/entities"
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	cards []*entities.Card
}

func New() *Deck {
	suits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	deck := make([]*entities.Card, 0)

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, &entities.Card{Suit: suit, Value: value})
		}
	}

	return &Deck{cards: deck}

}

func (d *Deck) String() string {
	return fmt.Sprintf("%s", d.cards)
}

func (d *Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.cards {
		newPosition := r.Intn(len(d.cards) - 1)

		d.cards[i], d.cards[newPosition] = d.cards[newPosition], d.cards[i]
	}
}
