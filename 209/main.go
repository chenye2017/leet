package main

import "fmt"

func main() {
	target := 11
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(target, nums))
}

/**
和通过叠加的方式很不错
 o(n) 时间复杂度
 o(1) 空间复杂度
*/
func minSubArrayLen(target int, nums []int) int {
	count := len(nums)

	start := 0
	end := 0

	sum := nums[start]

	min := 0
	noMin := true // 这块老师的方法就是取得整体size + 1, 然后末尾判断

	for {
		if sum >= target {
			leave := end - start + 1
			if noMin || min > leave {
				min = leave
				noMin = false
			}

			if end != start {
				sum = sum - nums[start]
			} else {
				sum = nums[end]
			}
			start++

			if start > end {
				break
			}

			continue
		} else {
			end++
			if end >= count {
				break
			}
			sum = sum + nums[end]
			continue
		}
	}

	return min
}
