package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card to be 'Two of Spades', instead got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Diamonds" {
		t.Errorf("Expected last card 'Ace of Diamonds', instead got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testingFile")

	d := newDeck()
	d.saveToFile("_testingFile")

	loadedD := newDeckFromFile("_testingFile")

	if len(loadedD) != 52 {
		t.Errorf("Expected 52 cards, instead got %v", len(loadedD))
	}

	os.Remove("_testingFile")
}
