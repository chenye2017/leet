package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	defer func() {
		fmt.Println(nums)
	}()

	start := 0
	for k, v := range nums {
		if k == 0 {
			start++
			continue
		} else {
			if v != nums[k-1] {
				nums[start] = v
				start++
			}
		}
	}
	nums = nums[0:start]
	return start
}
