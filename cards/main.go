package main

import "fmt"

func main() {

	cards := newDeck()
	// cards := deck{}
	// cards.print()

	// hand, cards := deal(cards, 10)
	// fmt.Println("Cards dealt")
	// cards.print()
	// fmt.Println("hand")
	// hand.print()
	cards.print()

	fmt.Println("shuffling...")
	cards.shuffle()

	cards.print()
}
