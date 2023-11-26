package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
	cards.shuffle()
	cards.printDeck()
	// newDeckFromFile("my_cards1.txt").printDeck()
	// hand, remainingDeck := deal(cards, 5)

	// hand.printDeck()
	// remainingDeck.printDeck()
}