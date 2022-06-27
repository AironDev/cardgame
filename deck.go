package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// define  deck type as a slice of strings
type deck []string

func newDeck() deck {
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
	return strings.Join(str, ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	deck := strings.Split(string(data), ",")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck
}

func (d deck) shuffleDeck() deck {
	rand.Seed(time.Now().UnixNano())

	//method 1
	//rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	//return d

	// method 2
	for i := range d {
		j := rand.Intn(len(d) - 1)
		d[j], d[i] = d[i], d[j]
	}

	return d
}
