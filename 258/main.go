package main

import "fmt"

func main() {
	num := 0
	fmt.Println(addDigits(num))
}

func addDigits(num int) int {
	total := 0
	for {
		tmp := num % 10
		total += tmp
		num /= 10
		if num == 0 {
			// 符合条件了
			if total < 10 {
				break
			} else {
				num = total
				total = 0
			}
		}
	}
	return total
}
