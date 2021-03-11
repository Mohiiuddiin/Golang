package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(gap(10, 300, 400)) //4,100,110
}

func gap(g, m, n int) []int {
	if g > (n - m) {
		return nil
	}
	var primeNums []int
	flag := 1
	for i := m; i <= n; i++ { //finding prime number

		if i != 2 && i%2 != 0 {
			flag = 1
			for j := 3; j <= int(math.Floor(math.Sqrt(float64(i)))); j += 2 {
				if i%2 == 0 {
					flag = 0
					break
				}
				if i%j == 0 {
					flag = 0
				}
			}

			if flag == 1 {
				primeNums = append(primeNums, i)
			}
		}

	}

	for i := 0; i < len(primeNums)-1; i++ {

		if (primeNums[i+1] - primeNums[i]) == g {
			return []int{primeNums[i], primeNums[i+1]}
		}

	}

	return nil
}
