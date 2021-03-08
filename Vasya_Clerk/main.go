package main

import (
	"fmt"
)

func main() {
	fmt.Println(tickets([]int{25, 25, 25, 25, 50, 50, 100}))
}

func tickets(peopleInLine []int) string {

	count1 := 0
	count2 := 0
	for i := 0; i < len(peopleInLine); i++ {
		if peopleInLine[i] == 25 {
			count1++
		} else if peopleInLine[i] == 50 {
			if count1 >= 25 {
				count1--
				count2++
			} else {
				return "NO"
			}
		} else if peopleInLine[i] == 100 {

			if count1 >= 1 && count2 >= 1 {
				count1--
				count2--
			} else if count1 >= 3 {
				count1 -= 3
			} else {
				return "NO"
			}
		} else {
			return "NO"
		}
	}

	return "YES"
}
