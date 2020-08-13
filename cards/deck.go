package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Create a new type of 'deck'
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//return two different slices
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	// []string(d) //convert to slice of string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// func newDeckFromFile(filename string) deck {
// 	bs, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		// Option 1 - log the error and default to newDeck()
// 		// Option 2 - log the error and quit program
// 		fmt.Println("Error:", err)

// 	}
// 	return deck(bs)
// }
