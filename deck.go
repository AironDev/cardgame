package main

import (
	"fmt"
	"os"
	"strings"
)

// define  deck type as a slice of strings
type deck []string

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, val := range cardValues {
			title := fmt.Sprintf("%s of %s", val, suite)
			cards = append(cards, title)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) toString() string {
	str := []string(d)
	return strings.Join(str, " ")
}

func writeToFile(data string) {
	os.WriteFile("deckfile", []byte(data), os.FileMode('w'))
}
