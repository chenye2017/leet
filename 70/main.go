package main

import "fmt"

func main() {
	n := 4
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		leave := i % 2
		if leave == 1 {
			first = first + second
		} else {
			second = first + second
		}
	}
	if (n % 2) == 1 {
		return first
	} else {
		return second
	}
}
