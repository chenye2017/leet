package main

import "fmt"

func main() {
	numbers := []int{-1, 0}
	target := -1

	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for {
		if start == end {
			return []int{}
		}

		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		}
		if sum > target {
			end--
		} else {
			start++
		}
	}
}
