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
// func NbDig(n int, d int) int {
//     if n == 0 {
//         return 0
//     }
    
//     count := 0
//     for i := 0; i <= n; i++ {
//         count += digits(i*i, d)
//     }
//     return count
// }

// func digits(val, d int) int {
//     if val == d {
//         return 1
//     }

//     count := 0
//     for val != 0 {
//         digit := val % 10
//         if digit == d {
//           count++
//         }
//         val /= 10
//     }
//     return count
// }
