package main

import "fmt"

func main() {
	n := 25
	fmt.Println(tribonacci(n))
}

func tribonacci(n int) int {
	t0 := 0
	t1 := 1
	t2 := 1

	switch n {
	case 0:
		return t0
	case 1:
		return t1
	case 2:
		return t2
	}

	for i := 3; i <= n; i++ {
		leave := i % 3
		if leave == 0 {
			t0 = t0 + t1 + t2
		} else if leave == 1 {
			t1 = t0 + t1 + t2
		} else if leave == 2 {
			t2 = t0 + t1 + t2
		}
	}

	switch n % 3 {
	case 0:
		return t0
	case 1:
		return t1
	case 2:
		return t2
	}
	return 0
}
