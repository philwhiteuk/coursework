package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	hand, cards := cards.deal(3)
	//hand.print()

	fmt.Println(hand.toByteSlice())
	//cards.print()
	saveToDisk(hand)
}
