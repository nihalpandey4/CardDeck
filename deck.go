package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize],d[handSize:];
}

func (d deck) toString ()  string {
	return strings.Join([]string(d),",")
}

func StringToDeck(deckString string) deck {
	return deck(strings.Split(deckString,","))
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()),0666);
}

func newDeckFromFile (fileName string) deck {
	fileData,error := os.ReadFile(fileName)
	if error !=nil {
		fmt.Println("Error:",error);
		os.Exit(1)
	}
	cardString := string(fileData)
	cards := StringToDeck(cardString);
	return cards;
}

func (d deck) shuffle() deck{
	for i:=0;i<len(d);i++{
		randomNumber := rand.Intn(len(d)-1)
		d[randomNumber],d[i] = d[i],d[randomNumber]
	}
	return d;
}