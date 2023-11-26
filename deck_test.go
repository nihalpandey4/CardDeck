package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	expectedLength := 52
	if len(cards) != expectedLength {
		t.Errorf("Expected deck length of %v , got %v",expectedLength, len(cards))
	}
}