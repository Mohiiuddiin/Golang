package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//creating a new type deck
//which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "1", "2", "3",
		"4", "5", "6", "7", "8", "9", "10"}

	//we will use underscore when we don't need the value of i and j
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
//converting deck to a slice of array []string(d) -->type conversion
//Joining all strings of a slice of array to a comma separated single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
//...
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//d means reciever,which is a copy of deck
func (d deck) print() {
	for i, card := range d {
		if i%2 == 0 {
			fmt.Println("")
		}

		fmt.Print(i, ". [", card, "] ")

	}
}
