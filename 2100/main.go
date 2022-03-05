package main

import "fmt"

func main() {
	security := []int{4, 3, 2, 1}
	time := 1

	fmt.Println(goodDaysToRobBank(security, time))
}

func goodDaysToRobBank(security []int, time int) []int {
	count := len(security)
	res := make([]int, 0)
	if time == 0 {
		for k := range security {
			res = append(res, k)
		}
		return res
	}

	if count <= 2*time {
		return []int{}
	}

	start := 0 + time       // [
	end := count - time - 1 // ]

	checkZ := func(i int, time int, arr []int) bool {
		tmpStart := i - time
		tmpEnd := i

		for {
			if tmpStart == i {
				break
			}

			if arr[tmpStart] >= arr[tmpStart+1] {
				tmpStart++
			} else {
				return false
			}
		}

		for {
			if tmpEnd == i+time {
				break
			}
			if arr[tmpEnd] <= arr[tmpEnd+1] {
				tmpEnd++
			} else {
				return false
			}
		}
		return true
	}

	for i := start; i <= end; i++ {
		if r := checkZ(i, time, security); r {
			res = append(res, i)
		}
	}
	return res
}
