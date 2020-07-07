package main

import "testing"

func TestInitializeUltimateDeck(t *testing.T) {

	var cards deck
	cards.initializeUltimateDeck()

	if len(cards) != 50 {
		t.Errorf("Error regarding the length of the deck")
	}

}
