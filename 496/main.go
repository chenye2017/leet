package main

import (
	"errors"
	"fmt"
)

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}

	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	type stack struct {
		values []int
	}
	push := func(s *stack, v int) {
		s.values = append(s.values, v)
	}
	isEmpty := func(s *stack) bool {
		if len(s.values) == 0 {
			return true
		} else {
			return false
		}
	}
	peak := func(s *stack) (int, error) {
		if empty := isEmpty(s); empty {
			return 0, errors.New("is empty")
		} else {
			return s.values[len(s.values)-1], nil
		}
	}
	pop := func(s *stack) {
		s.values = s.values[0 : len(s.values)-1]
		return
	}

	s := new(stack)

	m := make(map[int]int)

	index := 0
	count := len(nums2)
	for {
		if index >= count {
			break
		}

		v := nums2[index]

		if isEmpty(s) {
			push(s, v)
			index++
		} else {
			peakValue, _ := peak(s)
			if peakValue < v {
				m[peakValue] = v
				pop(s)
			} else {
				push(s, v)
				index++
			}
		}
	}

	for k, v := range nums1 {
		if tmp, ok := m[v]; ok {
			nums1[k] = tmp
		} else {
			nums1[k] = -1
		}
	}
	return nums1
}
