package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello world  "

	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")

	arr1 := make([]string, 0)
	for _, v := range arr {
		if len(v) > 0 {
			arr1 = append(arr1, v)
		}
	}

	count := len(arr1)
	for i := 0; i < count/2; i++ {
		arr1[i], arr1[count-1-i] = arr1[count-1-i], arr1[i]
	}

	return strings.Join(arr1, " ")
}
