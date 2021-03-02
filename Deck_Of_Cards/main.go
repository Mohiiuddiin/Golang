package main

import "fmt"

func main() {

	cards := newDeck()
	var choice string
	fmt.Println("----------Deck Of Cards----------")
	fmt.Println("\n1.Create and display all available cards.\n" +
		"2.Save all cards into a file.\n" +
		"3.Retrieve card imformation from a created file\n" +
		"4.Shuffle the existing Cards and Take a number of Cards from the shuffled bundle to see whats on your luck\n" +
		"5.Quit The Program\n\n")

	decision := "no"
	//if the value is no program will run repeatedly else if the value is yes then it will stop the program
	for decision == "no" {
		fmt.Println("\n\nEnter your choice(1/2/3/4/5) : ")
		fmt.Scan(&choice)
		if choice == "1" {
			println("\nAll Available Cards : ")
			cards.print()

		} else if choice == "2" {

			var fileName string
			fmt.Println("Please Enter the filename : ")
			fmt.Scan(&fileName)
			cards.saveToFile(fileName)
			fmt.Println("\nFile Saved Successfully[File Name : ", fileName+"]")

		} else if choice == "3" {

			var fileName string
			fmt.Println("\nPlease Enter the filename : ")
			fmt.Scan(&fileName)
			cards := newDeckFromFile(fileName)
			fmt.Println("\nAll information retrieved from [File Name : ", fileName+"]")
			cards.print()

		} else if choice == "4" {

			cards.shuffle()
			var num int
			fmt.Println("\n\nPlease enter the number of cards cards you want to take : ")
			fmt.Scan(&num)
			if num < 56 {
				hand, remainingCards := deal(cards, num)
				fmt.Println("\n\nCards On My Luck : ")
				hand.print()
				fmt.Println("\n\nRemaining Cards : ")
				remainingCards.print()
			} else {
				fmt.Println("\nYou entered an wrong number")
			}

		} else if choice == "5" {
			fmt.Println("\nDo you want to quit the program? reply (yes/no) : ")
			fmt.Scan(&decision)
			if decision == "yes" {
				fmt.Println("\n\n-----Thank You-----Thank You-----Thank You----- ")
				break
			}
		} else {
			fmt.Println("\nYour choice is wrong")
		}
	}

}
