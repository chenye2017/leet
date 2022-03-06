package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "Aabb"
	fmt.Println(frequencySort(s))
}

func frequencySort(s string) string {
	if s == "" {
		return ""
	}

	m := make(map[byte]int)

	// 每个单词的使用评率
	for _, v := range s {
		m[byte(v)]++
	}

	freq := make(map[int][]byte)
	for k, v := range m {
		freq[v] = append(freq[v], k)
	}

	// 数组存储使用评率
	arr := make([]int, 0)
	for k := range freq {
		arr = append(arr, k)
	}

	sort.SliceStable(arr, func(i, j int) bool {
		if arr[i] > arr[j] {
			return true
		} else {
			return false
		}
	})

	res := make([]byte, 0)
	for _, v := range arr { // v 是个数
		for _, v1 := range freq[v] { // v1数具体内容
			for i := 0; i < v; i++ {
				res = append(res, v1)
			}
		}
	}
	return string(res)
}
