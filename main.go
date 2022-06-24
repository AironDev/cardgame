package main

import "fmt"

func main() {
	var cards deck = NewDeck()
	//cards.print()
	hand, remCards := deal(cards, 5)

	hand.print()
	fmt.Println("hands above")
	remCards.print()
}
