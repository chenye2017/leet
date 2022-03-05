package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = nums2
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
	}
	// num2 接在前面
	if nums1[0] >= nums2[n-1] {
		// num1 直接移动
		for k, v := range nums2 {
			nums1[m+k] = nums1[k]
			nums1[k] = v
		}
	}
	// 正常合并
	start := 0
	for k, v := range nums2 {
		for {
			if v >= nums1[start] {
				start++
			} else {
				for i:=m+k;{
					nums1[m+k] = nums1[m]
				}
			}
		}
	}

}
