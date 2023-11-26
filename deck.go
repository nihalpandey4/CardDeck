package main

import "fmt"

// create a new type of deck, slice of strings
type deck []string

func newDeck() deck{
	cards := deck{}
	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"};
	cardTypes := []string{"Ace","King","Queen","Jack","Ten","Nine","Eight","Seven","Six","Five","Four","Three","Two"};

	for i:=0;i<len(cardSuits);i++ {
		for j:=0;j<len(cardTypes);j++ {
			cards = append(cards, (cardTypes[j]+" of "+cardSuits[i]));
		}
	}
	
	return cards;
}

func (d deck) printDeck(){
	for i:=0;i<len(d);i++ {
		fmt.Println(d[i]);
	}	
}