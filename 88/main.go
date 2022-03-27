package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 4, 5, 6, 0}
	nums2 := []int{3}
	merge(nums1, 5, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for k, v := range nums2 {
			nums1[k] = v
		}
		return
	}
	if n == 0 {
		return
	}

	// num2 直接接在后面
	if nums1[m-1] <= nums2[0] {
		for k, v := range nums2 {
			nums1[m+k] = v
		}
		return
	}
	// num2 接在前面
	if nums1[0] >= nums2[n-1] {
		// num1 直接移动 n 的位置
		for k, v := range nums2 {
			nums1[n+k] = nums1[k]
			nums1[k] = v
		}
		return
	}

	// num1 后移元素 自己的len
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}

	start := 0
	num1Index := n
	num2Index := 0

	for {
		if nums1[num1Index] > nums2[num2Index] {
			nums1[start] = nums2[num2Index]
			num2Index++
		} else {
			nums1[start] = nums1[num1Index]
			num1Index++
		}
		start++

		if start == m+n || num2Index == n {
			break
		}

		if num1Index == m+n {
			for i := num2Index; i < n; i++ {
				nums1[start] = nums2[i]
				start++
			}
			return
		}

	}
	return
}
