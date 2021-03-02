package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//converting deck to a slice of array []string(d) -->type conversion
//Joing all strings of a slice of array to a comma separated single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	//0666 means anyone can read or write this file[file permission]
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

func newDeckFromFile(filaName string) deck {
	byteSlice, err := ioutil.ReadFile(filaName)
	//way 1 : log the error and return a new deck()
	//way 2 : log the error and entirely quit the program with the error
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	sliceOfStrings := strings.Split(string(byteSlice), ",")

	return deck(sliceOfStrings)
}

//shuffle the deck of cards by swapping with random position
func (d deck) shuffle() {
	//creating new source to generate new random number each time we run..
	source := rand.NewSource(time.Now().UnixNano())
	randNum := rand.New(source)

	for index := range d {
		newPosition := randNum.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index] //swap
	}
}
