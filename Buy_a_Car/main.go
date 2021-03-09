package main

import "fmt"

func main() {
	fmt.Println()
}

func nbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {

	var num [2]int
	saved := 0.0
	lossByTwo := 0.5
	oldPrice := float64(startPriceOld)
	newPrice := float64(startPriceNew)
	savings := 0.0
	if startPriceOld >= startPriceNew {

		num[0], num[1] = 0, startPriceOld-startPriceNew

		return num
	} else {
		saved = savings + oldPrice
		countMonths := 0

		for saved < newPrice {
			countMonths++
			if countMonths%2 == 0 {
				percentLossByMonth += lossByTwo
			}
			savings += float64(savingperMonth)
			oldPrice = oldPrice - oldPrice*.01*percentLossByMonth
			newPrice = newPrice - newPrice*.01*percentLossByMonth

			saved = savings + oldPrice
		}
		num[0], num[1] = countMonths, int(saved-newPrice+0.5)
		return num
	}

}
