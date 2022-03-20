package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, -5, 6}
	fmt.Println(maxSubArrayC(arr))
}

// 升级存内存
func maxSubArrayC(nums []int) int {
	maxSum := nums[0]
	content := []int{nums[0]}
	nowSum := maxSum
	nowContent := make([]int, len(content))
	copy(nowContent, content)

	for k, v := range nums {
		if k == 0 {
			continue
		} else {
			if nowSum < 0 {
				nowSum = v
				nowContent = []int{v}
			} else {
				nowSum += v
				nowContent = append(nowContent, v)
			}
			if maxSum < nowSum {
				maxSum = nowSum
				content = make([]int, len(nowContent))
				copy(content, nowContent)

			}
		}
	}
	fmt.Println(content)
	return maxSum
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	nowSum := maxSum

	for k, v := range nums {
		if k == 0 {
			continue
		} else {
			if nowSum < 0 {
				nowSum = v
			} else {
				nowSum += v
			}
			if maxSum < nowSum {
				maxSum = nowSum
			}
		}
	}
	return maxSum
}
