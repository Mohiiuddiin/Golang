package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solequa(5))
}

//x2 - 4 * y2 = n
func solequa(n int) [][]int {

	r := int((math.Sqrt(float64(n)))) + 1
	in := [][]int{}
	for i := 1; i <= r; i++ {

		if n%i != 0 {
			continue
		}

		y := (n/i - i) / 4
		x := i + 2*y
		n1 := (x - 2*y) * (x + 2*y)
		if x >= 0 && y >= 0 && n1 == n {
			in = append(in, []int{x, y})
		}

	}

	return in
}

// Extended Euclidean Alogorithm for GCD
// func gcdExtended(a, b int) int {
// 	//base case
// 	if b == 0 {
// 		x, y = 1, 0
// 		return a
// 	} else {
// 		g := gcdExtended(b, a%b) // recursively find the gcd
// 		x1, y1 := x, y
// 		x = y1
// 		y = x1 - (a/b)*y1
// 		fmt.Println(x, y)
// 		return g
// 	}

// }
