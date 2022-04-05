package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 7))
}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := (start + end) / 2
		fmt.Println(start, end, mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if nums[start] > target {
		return start + 1
	} else {
		return start
	}

}

func searchInsert1(nums []int, target int) int {
	c := len(nums)
	index := sort.Search(c, func(i int) bool {
		return nums[i] >= target
	})
	return index
}
