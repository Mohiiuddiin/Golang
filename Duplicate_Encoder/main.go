package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(duplicateEncode("Mohiuddin"))
}

func duplicateEncode(word string) string {
	word = strings.ToLower(word)
	var str string
	for i := 0; i < len(word); i++ {
		if strings.Count(word, string(word[i])) > 1 {
			str = str + ")"
		} else {
			str = str + "("
		}
	}
	return str
}
