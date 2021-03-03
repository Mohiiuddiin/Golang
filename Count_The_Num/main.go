package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(nbDig(10, 1))
}

func nbDig(n int, d int) int {

	count := 0
	for i := 0; i <= n; i++ {
		//strconv is used to convert into basic data types or vice-versa
		count += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
	return count
}
