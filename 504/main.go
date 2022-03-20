package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := -8
	fmt.Println(convertToBase7(num))
}

/**
通用的进制相关
*/
func convertToBase7(num int) string {
	var lessZero bool
	if num < 0 {
		lessZero = true
		num *= -1
	}

	arr := make([]int, 0)

	for {
		leave := num % 7
		arr = append(arr, leave)
		num = num / 7
		if num == 0 {
			break
		}
	}

	//
	count := len(arr)
	var res string
	for i := count - 1; i >= 0; i-- {
		res += strconv.Itoa(arr[i])
	}

	if lessZero {
		res = "-" + res
	}

	return res
}
