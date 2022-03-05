package main

import "fmt"

func main() {
	s := []byte{byte('h'), byte('e'), byte('l'), byte('l'), byte('o')}
	reverseString(s)
}

func reverseString(s []byte) {
	defer func() {
		for _, v := range s {
			fmt.Println(string(v))
		}
	}()

	count := len(s)
	if count == 0 {
		return
	}

	start := 0
	end := count - 1
	for {
		if start >= end {
			return
		}
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
