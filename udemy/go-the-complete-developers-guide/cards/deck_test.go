package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected new deck to have length 52 but got %v", len(d))
	}

	if d[0] != "AC" {
		t.Errorf("Expected new deck to have first card 'AC' but got %v", d[0])
	}

	if d[len(d)-1] != "KS" {
		t.Errorf("Expected new deck to have last card 'KS' but got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	nd := newDeckFromFile("_decktesting")

	if len(d) != len(nd) {
		t.Errorf("Expected new deck from file to have same length as the deck it was created from, but got %v, when %v was expected", len(nd), len(d))
	}

	os.Remove("_decktesting")
}
