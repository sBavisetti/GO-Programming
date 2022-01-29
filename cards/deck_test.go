package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	//arrange
	d := newDeck()

	//assert
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52,  but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected First card in deck to be Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected Last card in deck to be King of Clubs but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	filename := "_deckTesting"
	os.Remove(filename)

	cardDeck := newDeck()
	cardDeck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != len(cardDeck) {
		t.Errorf("Expected loaded Card deck to have %v cards but found %v cards", len(cardDeck), len(loadedDeck))
	}

	os.Remove(filename)
}

func TestShuffle(t *testing.T) {

	filename := "_deckTesting"
	os.Remove(filename)

	cardDeck := newDeck()
	cardDeck.saveToFile(filename)
	cardDeck.shuffle()

	loadedDeck := newDeckFromFile(filename)

	var count int = 0

	for index := range cardDeck {

		if cardDeck[index] == loadedDeck[index] {
			count = count + 1
		}
	}

	if count == 52 {
		t.Errorf("original card deck and shuffled card deck are in the same order hence shuffling failed")
	}

	os.Remove(filename)
}

func TestDeal(t *testing.T) {

	cardDeck := newDeck()
	var handSize int = 10

	hand, remainingDeck := deal(cardDeck, handSize)

	if len(hand) != handSize {
		t.Errorf("Hand size expected to be %v but is %v", handSize, len(hand))
	}

	if hand[0] != cardDeck[0] {
		t.Errorf("Expected to find %v as the first card in hand but found %v", cardDeck[0], hand[0])

	}

	if hand[handSize-1] != cardDeck[handSize-1] {
		t.Errorf("Expected to find %v as the last card in hand but found %v", cardDeck[handSize-1], hand[handSize-1])

	}

	if remainingDeck[0] != cardDeck[handSize] {
		t.Errorf("Expected to find %v as the first card in remaining Deck but found %v", cardDeck[handSize], remainingDeck[0])
	}

	if remainingDeck[len(remainingDeck)-1] != cardDeck[len(cardDeck)-1] {
		t.Errorf("Expected to find %v as the last card in remaining Deck but found %v", cardDeck[len(cardDeck)-1], remainingDeck[len(remainingDeck)-1])
	}
}
