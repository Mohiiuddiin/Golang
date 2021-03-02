package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56 , but got length of %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card deck is Ace of Spades , but got card of deck is %v", d[0])
	}

	if d[len(d)-1] != "10 of Clubs" {
		t.Errorf("Expected first card of deck is 10 of Clubs , but got card of deck is %v", d[len(d)-1])
	}
}

func TestSavetoFileAndNewDeckFromFile(t *testing.T) {
	deck := newDeck()
	deck.saveToFile("_DeckTesting")
	loadedDeck := newDeckFromFile("_DeckTesting")

	if len(loadedDeck) != 56 {
		t.Errorf("Expecting 56 card of deck, but got %v card of deck", len(loadedDeck))
	}
}
