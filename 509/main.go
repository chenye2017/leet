package main

import "fmt"

func main() {
	n := 4
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	first := 0
	second := 1
	for i := 2; i <= n; i++ {
		tmp := second
		second = first + second
		first = tmp
	}
	return second
}
