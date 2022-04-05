package main

import (
	"fmt"
	"sort"
)

func main() {
	i := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(i))
}

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else {
			return false
		}
	})

	res := make([][]int, 0)
	for k, v := range intervals {
		if k == 0 {
			res = append(res, v)
		} else {
			count := len(res)
			lastEnd := res[count-1][1]
			nowStart := v[0]
			if lastEnd < nowStart {
				res = append(res, v)
			} else {
				if res[count-1][1] < v[1] {
					res[count-1] = []int{res[count-1][0], v[1]}
				}
			}
		}
	}
	return res
}
