package main

import "fmt"

func main() {
	x := 10
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	mod := 10
	yushu := 0
	arr := make([]int, 0)
	for {
		yushu = x % 10
		arr = append(arr, yushu)
		x = x / 10

		if x == 0 {
			break
		}
		mod *= 10
	}
	arrRevers := make([]int, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		arrRevers = append(arrRevers, arr[i])
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != arrRevers[i] {
			return false
		}
	}
	return true
}
