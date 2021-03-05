package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := "-1 -2 -3 -4 -5 -6 -7"
	fmt.Println(HighAndLow(in))
}

func HighAndLow(in string) string {
	var numarray []int
	var strSlice []string
	strSlice = strings.Split(in, " ")
	for i := 0; i < len(strSlice); i++ {

		num, _ := strconv.Atoi(string(strSlice[i]))
		numarray = append(numarray, num)

	}
	strSlice = nil
	numarray = sort(numarray)
	fmt.Println(numarray)
	strSlice = append(strSlice, strconv.Itoa(numarray[0]), strconv.Itoa(numarray[len(numarray)-1]))
	in = strings.Join(strSlice, " ")

	return in
}

func sort(numarray []int) []int {
	max := 0
	for i := 0; i < len(numarray); i++ {
		for j := i + 1; j < len(numarray); j++ {
			if numarray[j] > numarray[i] {
				max = numarray[j]
				numarray[j] = numarray[i]
				numarray[i] = max
			}
		}
	}
	return numarray
}
