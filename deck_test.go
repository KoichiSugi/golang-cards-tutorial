package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //*testing.T specifies a type of value
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first card of Ace of Spades, but got ", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Error("Expected first card of Ace of Spades, but got ", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting") //this mehtod returns slice of deck type
	if len(loadedDeck) != 16 {
		t.Error("Expected 16 cards in deck, but got ", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
