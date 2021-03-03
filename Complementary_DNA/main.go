package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "ATTGC"

	fmt.Println(dnaStrand(str))
}

func dnaStrand(dna string) string {

	var r rune
	var strSlice []string
	var tmpSlice []string
	var tmpStr string
	strSlice = strings.Split(dna, " ")
	k := len(strSlice)

	for i := 0; i < k; i++ {
		tmpStr = strSlice[i]
		for j := 0; j < len(strSlice[i]); j++ {
			if tmpStr[j] == 'C' {
				r = rune(tmpStr[j])
				r = 'G'
				tmpSlice = append(tmpSlice, string(r))
			} else if tmpStr[j] == 'G' {
				r = rune(tmpStr[j])
				r = 'C'
				tmpSlice = append(tmpSlice, string(r))
			} else if tmpStr[j] == 'A' {
				r = rune(tmpStr[j])
				r = 'T'
				tmpSlice = append(tmpSlice, string(r))
			} else if tmpStr[j] == 'T' {
				r = rune(tmpStr[j])
				r = 'A'
				tmpSlice = append(tmpSlice, string(r))
			} else {
				r = rune(tmpStr[j])
				tmpSlice = append(tmpSlice, string(r))
			}
		}
		tmpStr = strings.Join([]string(tmpSlice), "")
		strSlice[i] = tmpStr
		tmpSlice = nil
	}

	dna = strings.Join([]string(strSlice), " ")

	return dna

}
