package main

import "fmt"

func main() {
	s := "aabb"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	for k, v := range s {
		if m[v] == 1 {
			return k
		}
	}
	return -1
}
