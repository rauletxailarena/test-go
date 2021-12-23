package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	stringified_slice := strings.Join(strings_slice, ",")
	return stringified_slice
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		new_position := r.Intn(len(d) - 1)
		d[i], d[new_position] = d[new_position], d[i]
	}
}
