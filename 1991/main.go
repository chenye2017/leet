package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	nums = []int{1, 2, 3}
	nums = []int{2, 1, -1}
	fmt.Println(findMiddleIndex(nums))
}

func findMiddleIndex(nums []int) int {
	count := len(nums)
	if count == 1 {
		return 0
	}

	left := 0

	right := 0
	for k, v := range nums {
		if k == 0 {
			continue
		} else {
			right += v
		}
	}

	if left == right {
		return 0
	}

	for k, v := range nums {
		if k == 0 {
			continue
		} else {
			left += nums[k-1]
			right -= v
			if left == right {
				return k
			}
		}
	}
	return -1
}
