package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "aaa"
	s := "aa aa aa aa"

	fmt.Println(wordPattern(pattern, s))
}

func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")

	if len(pattern) != len(arr) {
		return false
	}

	m := make(map[string]byte)
	inUse := make(map[byte]struct{})

	for k, v := range arr {
		if _, ok := m[v]; !ok {
			if _, ok := inUse[pattern[k]]; ok {
				return false
			}
			m[v] = pattern[k]
			inUse[pattern[k]] = struct{}{}
		} else {
			if m[v] != pattern[k] {
				return false
			}
		}
	}

	return true
}
