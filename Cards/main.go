package main

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 6)
	// hand.print()
	// fmt.Println("Remaining Cards : ")
	// remainingCards.print()

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))

	//cards.print()
	// fmt.Println(cards.toString())

	cards.saveToFile("my_Cards")
}
