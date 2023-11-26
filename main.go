package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	hand.printDeck()
	fmt.Println("----------")
	remainingDeck.printDeck()
}