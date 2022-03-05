package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	defer func() {
		fmt.Println(nums)
	}()
	start := 0
	for _, v := range nums {
		if v != val {
			nums[start] = v
			start++
		}
	}
	nums = nums[0:start]
	return start
}
