package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// 'deck' is a slice of strings
type deck []string

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {

	for i, card := range d {

		fmt.Println(i, card)

	}

}

func deal(d deck, handSize int) (deck, deck) {

	hand := d[:handSize]
	remainingDeck := d[handSize:]
	return hand, remainingDeck
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return toDeck(string(content))
}

func (d deck) toString() string {

	resString := []string(d)
	return strings.Join(resString, ",")

}

func (d deck) shuffle() {

	deckSize := len(d) - 1
	var j int
	rand.Seed(time.Now().UnixNano())
	for i := range d {

		j = rand.Intn(deckSize)
		d[i], d[j] = d[j], d[i]
	}
}

func toDeck(data string) deck {

	return deck(strings.Split(data, ","))
}
