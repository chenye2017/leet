package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}

/**
这个老师的做法比我好，空间复杂度更低
他是利用map 中 key -> count, count -- 如果 count == 0 说明这个元素不存在了，不能匹配
*/
func intersect(nums1 []int, nums2 []int) []int {
	n1 := make(map[int]int)
	n2 := make(map[int]int)
	res := make(map[int]int)

	for _, v := range nums1 {
		if _, ok := n1[v]; ok {
			n1[v] += 1
		} else {
			n1[v] = 1
		}
	}

	for _, v := range nums2 {
		if _, ok := n2[v]; ok {
			n2[v] += 1
		} else {
			n2[v] = 1
		}
	}

	for k, v := range n1 {
		if c, ok := n2[k]; ok {
			var min int
			if v > c {
				min = c
			} else {
				min = v
			}
			res[k] = min
		}
	}

	resArr := make([]int, 0)
	for k, v := range res {
		for i := 0; i < v; i++ {
			resArr = append(resArr, k)
		}
	}
	return resArr
}
