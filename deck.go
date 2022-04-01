package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toStringFromDeck() string {
	stringSliceDeck := []string(d)
	stringDeck := strings.Join(stringSliceDeck, ",")
	return stringDeck
}

func toStringFromBytes(bytes []byte) string {
	return string(bytes)
}
func readFromFile(filename string) (deck, error) {
	bytes, error := ioutil.ReadFile(filename)
	cards := deck(strings.Split(toStringFromBytes(bytes), ","))
	return cards, error
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toStringFromDeck()), 0666)
}
