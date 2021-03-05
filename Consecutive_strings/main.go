package main

import (
	"fmt"
)

func main() {
	var strarr []string
	var str string
	//strarr = append(strarr, "zone", "abigail", "theta", "form", "libe", "zas")
	var k int

	fmt.Print("Enter the value of k : ")
	fmt.Scanln(&k)
	fmt.Println("Enter the values of string of array(write 'end' to stop taking more input) : ")
	for {
		fmt.Scanln(&str)
		if str == "end" {
			break
		} else {
			strarr = append(strarr, str)
		}
	}

	fmt.Println(longestConsec(strarr, k))
}

func longestConsec(strarr []string, k int) string {

	if k > 0 && k <= len(strarr) && strarr != nil {
		var tmpslice []string
		str := ""
		n := 0
		k1 := k - 1

		index := 0
		for i := 0; i < len(strarr)-k1; i++ {
			str = ""

			for j := 0; j <= k1; j++ {
				str = str + strarr[i+j]
			}
			tmpslice = append(tmpslice, str)

			if n < len(tmpslice[i]) {
				n = len(tmpslice[i])
				index = i
			}
		}
		return tmpslice[index]
	} else if strarr == nil {
		return ""
	} else {
		return ""
	}

}
