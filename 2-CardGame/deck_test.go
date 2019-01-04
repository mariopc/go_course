package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	testLen := 52
	if len(d) != testLen {
		t.Errorf("Expected deck length of %d, but got %d", testLen, len(d))
	}
	firstCard := d[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("Expected first card of \"Ace of Spades\" but got %s", d[0])
	}
	lastCard := d[len(d)-1]
	if lastCard != "King of Clubs" {
		t.Errorf("Expected first card of \"King of Clubs\" but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("__decktesting")

	deck := newDeck()
	deck.saveToFile("__decktesting")

	loadedDeck := newDeckFromFile("__decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("__decktesting")
}
