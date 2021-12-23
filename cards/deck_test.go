package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck lenght of 52, got %d", len(deck))
	}
	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of the deck to be Ace of spades, got %s", deck[0])
	}
	if deck[len(deck)-1] != "K of Clubs" {
		t.Errorf("Expected last card of the deck to be K of Clubs, got %s", deck[len(deck)-1])
	}
}

func TestSaveToToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck lenght of 52, got %d", len(deck))
	}

	os.Remove("_decktesting")

}
