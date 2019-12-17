package deckk

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deckk length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades at first position, but was %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deckManager := DeckManager{Data: newDeck()}
	deckManager.SaveToFile("_decktesting")

	loadedDeck, _ := NewDeckFromFile("_decktesting")

	if len(deckManager.Data) != len(loadedDeck.Data) {
		t.Errorf("Expected same langth %v %v", len(deckManager.Data), len(loadedDeck.Data))
	}
	os.Remove("_decktesting")
}
