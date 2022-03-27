package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else {
			if v-min > profit {
				profit = v - min
			}
		}
	}

	return profit
}
