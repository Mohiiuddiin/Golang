package main

import "fmt"

func main() {

	cards := newDeck()
	var choice string
	fmt.Println("1. Create and display all available cards.\n" +
		"2.Save all cards into a file.\n" +
		"3.Retrieve card imformation from a created file\n" +
		"4.Shuffle the existing Cards and Take a number of Cards from the shuffled bundle to see whats on your luck\n")
	fmt.Println("Enter your choice(1/2/3/4/5) : ")
	fmt.Scanln(&choice)

	if choice == "1" {

		cards.print()
	} else if choice == "2" {

		//cards := newDeck()
		var fileName string
		fmt.Scanln(&fileName)
		cards.saveToFile(fileName)

	} else if choice == "3" {

		var fileName string
		fmt.Scanln(&fileName)
		cards := newDeckFromFile(fileName)
		cards.print()

	} else if choice == "4" {
		//cards := newDeck()
		cards.shuffle()
		//cards.print()
		var num int
		fmt.Println("\nPlease enter the number of cards cards you want to take : ")
		fmt.Scan(&num)
		if num < 56 {
			hand, remainingCards := deal(cards, num)
			hand.print()
			fmt.Println("\n\nRemaining Cards : ")
			remainingCards.print()
		} else {
			fmt.Println("You entered an wrong number")
		}
	} else {
		fmt.Println("Your choice is wrong")
	}

}
