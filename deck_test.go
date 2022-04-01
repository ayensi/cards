package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf(fmt.Sprintf("Expected value is 52 but got %v instead", len(d)))
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected value is King of Clubs but got %v instead", d[len(d)-1])
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected value is Ace of Spades but got %v instead", d[0])
	}
}

func TestSaveDeckToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	file, err := os.Open("_decktesting") // For read access.
	if err != nil {
		t.Errorf(err.Error())
	}
	file.Close()

	newDeckFromFile, error := readFromFile("_decktesting")

	if error != nil {
		t.Errorf(err.Error())
	}

	if len(newDeckFromFile) != 52 {
		t.Errorf(fmt.Sprintf("Expected value is 52 but got %v instead", len(newDeckFromFile)))
	}

}
