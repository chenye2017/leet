package main

import "fmt"

func main() {
	n := 5
	fmt.Println(generate(n))
}

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	start := [][]int{{1}, {1, 1}}

	for i := 3; i <= numRows; i++ {
		tmp := []int{1}

		last := start[i-2]

		count := len(last)
		for i := 1; i < count; i++ {
			tmp = append(tmp, last[i]+last[i-1])
		}
		tmp = append(tmp, 1)
		start = append(start, tmp)
	}

	return start
}
