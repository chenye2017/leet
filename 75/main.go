package main

import "fmt"

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
}

func sortColors(nums []int) {
	defer func() {
		fmt.Println(nums)
	}()

	zero := 0
	two := len(nums) - 1

	for k, v := range nums {
		if v == 0 {
			if k != zero {
				nums[k], nums[zero] = nums[zero], nums[k]
			}
			zero++
		} else if v == 2 {
			for {
				if nums[two] != 2 {
					break
				} else {
					two--
					if two <= k {
						break
					}
				}
			}
			if two <= k {
				return
			}
			if k != two {
				nums[k], nums[two] = nums[two], nums[k]
			}
			two--
			// nums[k] 可能是 0
			if nums[k] == 0 {
				if k != zero {
					nums[k], nums[zero] = nums[zero], nums[k]
				}
				zero++
			}
		}
	}

	return
}

func sortColors1(nums []int) {
	defer func() {
		fmt.Println(nums)
	}()
	start := 0

	// 0 元素提取道最前面
	for k, v := range nums {
		if v == 0 {
			if k != start {
				nums[start], nums[k] = nums[k], nums[start]
			}
			start++
		}
	}

	// 0 个元素是 0~start-1
	for k, v := range nums {
		if k < start {
			continue
		} else {
			if v == 1 {
				if k != start {
					nums[start], nums[k] = nums[k], nums[start]
				}
				start++
			}
		}
	}

	return
}
