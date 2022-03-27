package main

import "fmt"

func main() {
	arr := []int{1, 3, 2}
	fmt.Println(isMonotonic(arr))
}

func isMonotonic(nums []int) bool {
	count := len(nums)
	if count <= 1 {
		return true
	}
	var flag int
	var trend bool
	for i := 1; i < count; i++ {
		if nums[i] == nums[i-1] {
			continue
		} else {
			tmp := nums[i] > nums[i-1]
			if flag == 0 {
				trend = tmp
				flag = 1
			} else {
				if tmp != trend {
					return false
				}
			}
		}
	}
	fmt.Println(trend)

	return true
}
