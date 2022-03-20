package main

import "fmt"

func main() {
	arr := []int{100, 1, 11, 1, 120, 111, 123, 1, -1, -100}
	fmt.Println(nextGreaterElements(arr))
}

func nextGreaterElements(nums []int) []int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	count := len(nums) * 2
	nums = append(nums, nums...)

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
	peak := func(s *stack) int {
		return s.values[len(s.values)-1]
	}
	pop := func(s *stack) {
		s.values = s.values[0 : len(s.values)-1]
		return
	}

	startIndex := 0
	s := new(stack)
	m := make(map[int][]int)

	for {
		if startIndex >= count {
			break
		}
		v := nums[startIndex]
		if isEmpty(s) {
			push(s, v)
			startIndex++
		} else {
			peak := peak(s)
			if peak >= v {
				push(s, v)
				startIndex++
			} else {
				if len(m) > count {
					break
				} else {
					if _, ok := m[peak]; !ok {
						m[peak] = make([]int, 0)
					}
					m[peak] = append(m[peak], v)
					pop(s)
				}
			}
		}
	}

	for k, v := range tmp {
		if vv, ok := m[v]; ok {
			tmp[k] = vv[0]
			m[v] = m[v][1:]
		} else {
			tmp[k] = -1
		}
	}
	return tmp
}
