package main

import (
	"fmt"
)

/*
 * Complete the 'minimumLoss' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts LONG_INTEGER_ARRAY price as parameter.
 */

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func minimumLoss(price []int64) int64 {
	var min_loss int64
	min_loss = 100
	a := price
	for _, x := range price {
		x, a = a[0], a[1:]
		for _, y := range a {
			if x > y {
				min_loss = min(min_loss, x-y)
			}
		}
	}
	return min_loss
}

func main() {

	price := []int64{20, 7, 8, 2, 5}

	result := minimumLoss(price)

	fmt.Println(result)
}
