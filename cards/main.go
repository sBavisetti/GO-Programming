package main

import (
	"fmt"
	"os"
)

func main() {

	// Deleting anyFiles with the name "_savedDeck"
	fileName := "_savedDeck"
	os.Remove(fileName)

	// creating a newDeck and printing the results
	cards := newDeck()
	fmt.Println("Printing the results of New Deck")
	cards.print()

	// creating a newDeck and saving it to a file
	// and retriving the results back
	newCards := newDeck()
	fmt.Println("Printing the New Deck")
	newCards.print()
	newCards.saveToFile(fileName)
	loadDeck := newDeckFromFile(fileName)
	fmt.Println("Printing the results of Loaded Deck from ", fileName)
	loadDeck.print()

	// shuffling a newDeck and Dealing
	// shuffling
	newCardsDeck := newDeck()
	fmt.Println("Before Shuffling : ")
	newCardsDeck.print()
	newCardsDeck.shuffle()
	fmt.Println("After Shuffling : ")
	newCardsDeck.print()

	// dealing
	hand, remainingDeck := deal(newCardsDeck, 5)
	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remaining Deck: ")
	remainingDeck.print()

	// deleting the savedDeck file
	os.Remove(fileName)

}
