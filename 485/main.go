package main

import "fmt"

func main() {
	nums := []int{0}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	// [1,1,0,1,1,1]

	max := 0
	start := 0
	end := 0
	noStart := true
	for k, v := range nums {
		if v == 1 {
			end = k
			if noStart {
				start = k
				noStart = false
			}
		} else {
			if !noStart {
				tmp := end - start + 1
				if tmp > max {
					max = tmp
				}
				noStart = true
			}
		}
	}

	if !noStart {
		tmp := end - start + 1
		if tmp > max {
			return tmp
		}
	}

	return max
}
