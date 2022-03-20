package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr, 4, -1))
}

func findClosestElements(arr []int, k int, x int) []int {
	po := sort.SearchInts(arr, x)

	// po 把元素切成2份
	left := make([]int, po)
	right := make([]int, len(arr)-po)

	copy(left, arr[0:po])
	copy(right, arr[po:])

	leftIndex := po - 1
	rightIndex := 0

	res := make([]int, 0)

	for {
		var leftSub int
		var rightSub int

		if leftIndex < 0 {
			leftSub = -1
		} else {
			leftSub = int(math.Abs(float64(left[leftIndex] - x)))
		}
		if rightIndex >= len(arr)-po {
			rightSub = -1
		} else {
			rightSub = int(math.Abs(float64(right[rightIndex] - x)))
		}

		if leftSub == -1 && rightSub == -1 {
			break
		}

		if rightSub == -1 || (leftSub <= rightSub && leftSub != -1) {
			res = append(res, left[leftIndex])
			leftIndex--
		}
		if leftSub == -1 || (leftSub > rightSub && rightSub != -1) {
			res = append(res, right[rightIndex])
			rightIndex++
		}

		if len(res) >= k {
			break
		}
	}

	sort.SliceStable(res, func(i, j int) bool {
		if res[i] < res[j] {
			return true
		} else {
			return false
		}
	})
	return res
}
