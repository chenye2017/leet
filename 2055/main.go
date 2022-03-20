package main

import "fmt"

func main() {
	s := "***|**|*****|**||**|*"
	q := [][]int{{4, 5}, {14, 17}, {5, 11}, {15, 16}}

	/*s = "**|**|***|"
	q = [][]int{{2, 5}, {5, 9}}*/

	s = "*|*||||**|||||||*||*||*||**|*|*||*"
	q = [][]int{{0, 33}}

	fmt.Println(platesBetweenCandles(s, q))
}

func platesBetweenCandles(s string, queries [][]int) []int {
	m := make([]int, 0)

	// 获取到所有位置
	for k, v := range s {
		if v == '|' {
			m = append(m, k)
		}
	}

	res := make([]int, 0)

	for _, v := range queries {
		start := -1
		startIndex := -1
		end := -1
		endIndex := -1

		// 找出开始索引和结束索引
		for k, v1 := range m {
			if v1 >= v[0] && start == -1 {
				// 第一次赋值
				start = v1
				startIndex = k
				end = v1
				endIndex = k
			}
			if v[1] >= v1 {
				end = v1
				endIndex = k
			}
		}

		fmt.Println(start, end, startIndex, endIndex, m, "----")
		// 没有边界
		if start == -1 || end == -1 || start >= end {
			res = append(res, 0)
		} else {
			inum := endIndex - startIndex + 1 // 竖线的个数 1, 2, 3, 4
			totalnum := end - start + 1       // 总的个数 1, 2, 3
			res = append(res, totalnum-inum)
		}

	}
	return res
}

func platesBetweenCandles1(s string, queries [][]int) []int {

	arr := make([]int, 0)
	var irune = '|'
	var xrune = '*'

	for _, v := range queries {
		tmp := s[v[0] : v[1]+1]

		start := -1
		have := 0
		realHave := 0
		for k, v := range tmp {
			if v == irune {
				if start == -1 {
					start = k
				} else {
					realHave = have
				}
			}
			if v == xrune {
				if start > -1 {
					have++
				}
			}
		}
		arr = append(arr, realHave)
	}
	return arr
}
