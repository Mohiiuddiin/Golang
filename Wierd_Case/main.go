package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var str string
	str = toWeirdCase("This is a test looks like you passed")
	fmt.Println("\n" + str)
}
func toWeirdCase(str string) string {

	var r rune
	var strSlice []string
	var tmpSlice []string
	var tmpStr string
	strSlice = strings.Split(str, " ")
	k := len(strSlice)

	for i := 0; i < k; i++ {
		tmpStr = strSlice[i]
		for j := 0; j < len(strSlice[i]); j++ {
			if j%2 == 0 {
				r = rune(tmpStr[j])
				r = unicode.ToUpper(r)
				tmpSlice = append(tmpSlice, string(r))
			} else {
				r = rune(tmpStr[j])
				r = unicode.ToLower(r)
				tmpSlice = append(tmpSlice, string(r))
			}
		}
		tmpStr = strings.Join([]string(tmpSlice), "")
		strSlice[i] = tmpStr
		tmpSlice = nil
	}

	str = strings.Join([]string(strSlice), " ")

	return str
}
