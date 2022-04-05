package main

import "fmt"

func main() {
	s := "cupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupucu"
	//s = "aba"
	//s = "abc"

	fmt.Println(validPalindrome(s))
}

func validPalindrome(s string) bool {
	// 转数组，双指针移动，最多一次机会
	arr := []byte(s)

	start := 0
	end := len(arr) - 1
	chance := 0
	shan := 0
	shanStart := 0
	shanEnd := 0

	for start < end {
		if arr[start] == arr[end] {
		} else {
			if chance < 1 {
				shanStart = start
				shanEnd = end
				start++
				chance++
				shan++
				continue
			} else {
				if shan < 2 {
					start = shanStart
					end = shanEnd - 1
					shan++
					continue
				} else {
					return false
				}
			}
		}
		start++
		end--
	}
	return true
}
