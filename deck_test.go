package main

import "testing"

func TestInitializeUltimateDeck(t *testing.T) {

	d := initializeUltimateDeck()

	if len(d) != 90 {
		t.Errorf("Error regarding the length of the deck")
	}

}
