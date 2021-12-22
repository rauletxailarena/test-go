package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{
		"Spades", "Hearts", "Diamonds", "Clubs",
	}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Q", "K",
	}
	my_deck := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			my_deck = append(my_deck, value+" of "+suit)
		}
	}
	return my_deck
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	strings_slice := []string(d)
	stringified_slice := strings.Join(strings_slice, ", ")
	return stringified_slice
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
