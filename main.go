package main

import (
	"cards_objects/deck"
	"fmt"
)

func main() {

	d := deck.New()
	fmt.Printf("%s\n", d)
	d.Shuffle()
	fmt.Printf("%s", d)

}
