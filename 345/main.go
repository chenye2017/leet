package main

import "fmt"

func main() {
	s := "leetcode"
	fmt.Println(reverseVowels(s))
}

func reverseVowels(s string) string {
	count := len(s)

	if count == 0 {
		return ""
	}

	arr := []byte(s)

	start := 0
	end := count - 1

	m := map[byte]struct{}{
		'a': {},
		'A': {},
		'e': {},
		'E': {},
		'i': {},
		'I': {},
		'o': {},
		'O': {},
		'u': {},
		'U': {},
	}

	for {
		if start >= end {
			return string(arr)
		}
		if _, ok := m[arr[start]]; !ok {
			start++
			continue
		}
		if _, ok := m[arr[end]]; !ok {
			end--
			continue
		}
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
