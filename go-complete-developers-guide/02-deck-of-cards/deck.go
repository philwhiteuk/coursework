package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	newDeck := deck{}
	suits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			newDeck = append(newDeck, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return newDeck
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func saveToDisk(d deck) {
	ioutil.WriteFile("cards", d.toByteSlice(), os.FileMode(0555))
}
