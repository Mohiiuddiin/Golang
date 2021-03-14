package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ListSquared(1, 250))
}

func ListSquared(m, n int) [][]int {
	sum := 0
	in := [][]int{}

	for i := m; i <= n; i++ {
		for j := 1; j <= int(i/2)+1; j++ {
			if i%j == 0 {
				sum += j * j
			} else if j == int(i/2)+1 {
				sum += i * i
				//fmt.Println(i)
			}

		}
		_, frac := math.Modf(math.Sqrt(float64(sum)))
		if frac == 0.0 {
			in = append(in, []int{i, sum})
		}
		sum = 0

	}
	return in
}
