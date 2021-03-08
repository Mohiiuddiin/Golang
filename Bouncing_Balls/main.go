package main

import "fmt"

func main() {
	fmt.Println(bouncingBall(3, 0.66, 1.5))
}

func bouncingBall(h, bounce, window float64) int {
	if h > 0 && bounce > 0 && bounce < 1 && window < h {

		count := 1 //when falling 1st time she will see it once
		eachBounce := h * bounce

		for eachBounce >= window {
			eachBounce = eachBounce * bounce
			count = count + 2 //for up + down it will be inc by 2
		}

		return count
	} else {
		return -1
	}
}
