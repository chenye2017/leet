package main

import "fmt"

func main() {

	nums := []int{0}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	defer func() {
		fmt.Println(nums)
	}()

	start := 2

	if len(nums) <= 2 {
		return 2
	}

	for k, v := range nums {
		if k < 2 {
			continue
		} else {
			if v == nums[start-2] {
				continue
			} else {
				nums[start] = v
				start++
			}
		}
	}

	nums = nums[0:start]
	return start
}
