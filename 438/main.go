package main

import "fmt"

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}

/**
我们在做m[byte]int 这种 字符-》position 索引的时候一定要想着要不要做，因为字符可能重复
所以我们是否只需要记住字符的数量就好了，如果我们不在乎截取的字符的话
*/
func findAnagrams(s string, p string) []int {
	cp := len(p)
	cs := len(s)

	arr := []byte(s)

	res := make([]int, 0)
	if cp > cs {
		return res
	}

	start := 0
	end := cp - 1

	tmp := arr[start : end+1]

	// 构建
	m := make(map[byte]int)
	for _, v := range tmp {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	target := make(map[byte]int)
	for _, v := range []byte(p) {
		if _, ok := target[v]; ok {
			target[v] += 1
		} else {
			target[v] = 1
		}
	}

	// 比较2个map是否相等
	compare := func(m, target map[byte]int) bool {
		for k, v := range target {
			if tmp := m[k]; tmp != v {
				return false
			}
		}
		return true
	}

	if compare(m, target) {
		res = append(res, 0)
	}
	for {
		m[arr[start]] -= 1
		// fmt.Println(string(arr[start]), m, "sss")
		end++
		start++
		if end >= cs {
			return res
		}
		if _, ok := m[arr[end]]; ok {
			m[arr[end]] += 1
		} else {

			m[arr[end]] = 1
		}

		// fmt.Println(m, "----")
		if compare(m, target) {
			res = append(res, start)
		}
	}
}
