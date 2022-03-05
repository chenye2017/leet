package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9

	arr = []int{3, 2, 4}
	target = 6

	arr = []int{3, 3}
	target = 6

	a := twoSum(arr, target)
	fmt.Println(a)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int][]int)

	for k, v := range nums {
		if m[v] == nil {
			m[v] = make([]int, 0)
		}
		m[v] = append(m[v], k)
	}

	for k, v := range m {
		sub := target - k
		if position, ok := m[sub]; ok {
			if sub == k {
				if len(position) > 1 {
					return []int{position[0], position[1]}
				}
			} else {
				return []int{v[0], position[0]}
			}
		}
	}
	return []int{}
}
