package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(nums))
}

func findUnsortedSubarray(nums []int) int {
	nums1 := make([]int, len(nums))

	copy(nums1, nums)

	sort.SliceStable(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		} else {
			return false
		}
	})

	left := -1
	right := -1
	for k, v := range nums {
		if v != nums1[k] {
			left = k
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != nums1[i] {
			right = i
			break
		}
	}

	if left == -1 || right == -1 {
		return 0
	}
	return right - left + 1
}
