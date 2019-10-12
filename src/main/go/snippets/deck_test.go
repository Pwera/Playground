package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades at first position, but was %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck, _ := newDeckFromFile("_decktesting")

	if len(d) != len(loadedDeck) {
		t.Errorf("Expected same langth %v %v", len(d), len(loadedDeck))
	}
	os.Remove("_decktesting")
}
