package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"

	haystack = "aaaaa"
	needle = "bba"

	haystack = ""
	needle = ""

	haystack = "mississippi"
	needle = "issip"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}

	haystackArr := []rune(haystack)
	needleArr := []rune(needle)

	if len(haystackArr) < len(needleArr) {
		return -1
	}

	start := 0
	end := start + len(needleArr)
	for {
		if needle == string(haystackArr[start:end]) {
			return start
		}
		start += 1
		end = start + len(needleArr)
		if end > len(haystackArr) {
			return -1
		}
	}
}

func strStr1(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}

	m := make(map[rune]int)

	for k, v := range []rune(haystack) {
		if _, ok := m[v]; !ok {
			// 第一次记录
			m[v] = k
		}
	}
	needleArr := []rune(needle)
	firstSearch := needleArr[0]
	firstSearchPosition, ok := m[firstSearch]
	if !ok {
		return -1
	}

	haystackArr := []rune(haystack)
	if firstSearchPosition+len(needleArr) > len(haystack) {
		return -1
	}

	if needle != string(haystackArr[firstSearchPosition:firstSearchPosition+len(needleArr)]) {
		return -1
	}
	return firstSearchPosition
}
