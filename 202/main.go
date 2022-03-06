package main

import "fmt"

func main() {
	n := 2
	fmt.Println(isHappy(n))
}

func isHappy(n int) bool {
	m := map[int]struct{}{n: {}}

	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	gen := func(number int) int {
		arr := make([]int, 0)
		for {
			if number == 0 {
				break
			}
			tmp := number % 10
			arr = append(arr, tmp)
			number /= 10
		}
		sum := 0
		for _, v := range arr {
			sum += v * v
		}
		return sum
	}

	for {
		if n = gen(n); n == 1 {
			return true
		} else {
			if _, ok := m[n]; ok {
				return false
			} else {
				m[n] = struct{}{}
			}
		}
	}
}
