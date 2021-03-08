package main

import (
	"fmt"
)

func main() {
	fmt.Println(revrot("733049910872815764", 5)) //330479108928157
}

// func rev_rot(s string, n int) string {
// 	sum1 := 0
// 	sum2 := 0
// 	num := 0
// 	var str []string
// 	if len(s) < n || s == "" || n == 0 {
// 		return ""
// 	} else {
// 		tmp := ""
// 		for i := 0; i < len(s); i++ {
// 			if
// 		}
// 		fmt.Println(str)
// 		for i := 0; i < len(str); i++ {
// 			if i < n {
// 				num, _ = strconv.Atoi(str[i])
// 				sum1 += num * num * num
// 			} else if i >= n {
// 				num, _ = strconv.Atoi(str[i])
// 				sum2 += num * num * num
// 			}

// 		}
// 		fmt.Println(str, sum1, sum2) //1

// 		if sum1%2 == 0 {
// 			for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
// 				str[i], str[j] = str[j], str[i]
// 			}
// 		} else {
// 			tmp = str[0]
// 			for i := 0; i < n; i++ {
// 				str[i] = str[i+1]
// 			}
// 			str[n-1] = tmp
// 		}

// 		if sum2%2 == 0 {
// 			for i, j := n, len(str)-1; i < j; i, j = i+1, j-1 {
// 				str[i], str[j] = str[j], str[i]
// 			}
// 		} else { //working
// 			tmp = str[n]
// 			for i := n; i < (len(str) - 1); i++ {
// 				str[i] = str[i+1]
// 			}
// 			str[len(str)-1] = tmp
// 		}

// 		return strings.Join(str, "")
// 	}
// }
//

func revrot(s string, n int) string {
	if len(s) < n || n == 0 {
		return ""
	}
	str := ""
	for i := 0; i < len(s)/n; i++ {
		substr := s[0+n*i : n*(i+1)]
		sum := 0
		for _, r := range substr { //r is a rune type of value
			sum += int((r - 48) * (r - 48) * (r - 48))
		}
		if sum%2 == 0 {
			s2 := ""
			//reverse
			for i := len(substr); i > 0; i-- {
				s2 += string(substr[i-1])
			}
			str += s2
		} else {
			//rotate
			str += substr[1:] + substr[:1]
		}
	}
	return str
}
