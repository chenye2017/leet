package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]struct{})
	res := make(map[int]struct{})

	for _, v := range nums1 {
		m1[v] = struct{}{}
	}

	for _, v := range nums2 {
		if _, ok := m1[v]; ok {
			res[v] = struct{}{}
		}
	}

	arr := make([]int, 0)
	for k := range res {
		arr = append(arr, k)
	}
	return arr
}
