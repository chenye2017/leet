package main

import "fmt"

func main() {
	security := []int{1, 2, 3, 4, 5, 6}
	time := 2

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

	if count < (2*time + 1) {
		return []int{}
	}

	// 穷举吧
	start := 0 + time       // 闭区间 2
	end := count - 1 - time //

	for i := start; i <= end; i++ {
		fmt.Println(i, "=====")
		flag := true
		// i 就是那个连续数
		for i1 := i - time; i1 < i; i1++ {
			if security[i1] < security[i1+1] {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}

		for i2 := i; i2 < i+time; i2++ {
			if security[i2] > security[i2+1] {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}

	return res
}
