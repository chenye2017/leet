package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes2(nums)
}

/**
老师的方法，把有数据的元素前置，剩下的都赋值 0
*/
func moveZeroes2(nums []int) {
	defer func() {
		fmt.Println(nums)
	}()
	start := 0
	for _, v := range nums {
		if v != 0 {
			nums[start] = v
			start++
		}
	}
	for i := start; i < len(nums); i++ {
		nums[i] = 0
	}
}

/**
保证了顺序
*/
func moveZeroes(nums []int) {
	defer func() {
		fmt.Println(nums)
	}()

	start := 0
	count := len(nums)
	for {
		if start >= count {
			return
		}

		if nums[start] == 0 {
			next := start + 1
			for {
				if next >= count {
					return
				}

				if nums[next] == 0 {
					next++
					continue
				} else {
					nums[start], nums[next] = nums[next], nums[start]
					start++
					break
				}
			}

		} else {
			start++
			continue
		}
	}
}

/**
没有保证顺序
*/
func moveZeroes1(nums []int) {
	defer func() {
		fmt.Println(nums)
	}()

	if len(nums) == 0 {
		return
	}
	start := 0
	end := len(nums) - 1

	for {
		if start >= end {
			return
		}
		if nums[start] == 0 {
			if nums[end] != 0 {
				nums[start], nums[end] = nums[end], nums[start]
				start++
				end--
			} else {
				end--
				continue
			}
		} else {
			start++
		}
	}
}
