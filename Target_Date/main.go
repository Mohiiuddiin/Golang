package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(dateNbDays(100, 150, 2))
}

func dateNbDays(a0 int, a int, p int) string {
	day := 0
	n1, n2 := 0.0, 0.0
	n1 = float64(a0)
	n2 = float64(a)
	for n1 <= n2 {
		day++
		n1 += float64(n1) * float64(p) / 36000
	}
	input := "2016-01-01"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	// t := time.Date(2016, 01, 01, 0, 0, 0, 0, time.Local)
	// fmt.Println(t)
	t2 := t.AddDate(0, 0, day)

	return string(t2.Format("2006-01-02"))
}
