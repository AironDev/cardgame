package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffleDeck()
	cards.print()
}
