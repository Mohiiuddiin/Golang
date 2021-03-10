package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solution([]int{10, 18, 19, 20, 21, 27, 28, 36, 37, 38, 39, 40, 44, 49, 50, 51, 52, 59}))
}

func solution(list []int) string {
	start, end := 0, 0
	var s string
	var intSlices []string

	for i := 0; i < (len(list) - 2); i++ {

		if list[i+2] == (list[i]+2) && list[i+1] == (list[i]+1) {
			start = list[i]
			for j := i + 1; j < len(list); j++ {
				for list[j] == (list[i] + 1) {
					end = list[j]
					i++
				}
			}
			s = strconv.Itoa(start) + "-" + strconv.Itoa(end)
		} else {
			s = strconv.Itoa(list[i])
		}
		intSlices = append(intSlices, s)
		if i == (len(list)-3) || i == (len(list)-2) {

			k := 0
			for k < len(list)-i {
				i++
				s = strconv.Itoa(list[i])
				intSlices = append(intSlices, s)
				k++
			}
		}
	}
	return strings.Join(intSlices, ",")
}
