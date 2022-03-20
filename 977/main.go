package main

import "fmt"

func main() {
	nums := []int{-4, -1}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	count := len(nums)

	if count == 0 {
		return []int{}
	}

	first := nums[0]

	// 所有元素都大于0
	if first >= 0 {
		for k, v := range nums {
			nums[k] = v * v
		}
		return nums
	}

	// 有元素小于0
	less := []int{}
	for _, v := range nums {
		if v < 0 {
			less = append(less, v*-1)
		}
	}

	lessStart := len(less) - 1

	res := []int{}
	start := 0

	for {
		if start > count-1 {
			break
		}

		if nums[start] < 0 {
			start++
			continue
		}

		if lessStart >= 0 && nums[start] > less[lessStart] {
			res = append(res, less[lessStart]*less[lessStart])
			lessStart -= 1
		} else {
			res = append(res, nums[start])
			start++
		}

	}

	if lessStart != -1 {
		for i := lessStart; i > -1; i-- {
			res = append(res, less[i]*less[i])
		}
	}

	return res
}
