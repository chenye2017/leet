package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}
	cost = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	daoda := len(cost)

	daodaArr := []int{0, 0}

	for k := range cost {
		if k < 2 {
			continue
		} else {
			tmp := min(daodaArr[k-2]+cost[k-2], daodaArr[k-1]+cost[k-1])
			daodaArr = append(daodaArr, tmp)
		}
	}

	// 最后一步单独算
	return min(daodaArr[daoda-1]+cost[daoda-1], daodaArr[daoda-2]+cost[daoda-2])
}
