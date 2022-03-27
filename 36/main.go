package main

import (
	"fmt"
)

func main() {
	arr := [][]string{
		{".", "8", "7", "6", "5", "4", "3", "2", "1"},
		{"2", ".", ".", ".", ".", ".", ".", ".", "."},
		{"3", ".", ".", ".", ".", ".", ".", ".", "."},
		{"4", ".", ".", ".", ".", ".", ".", ".", "."},
		{"5", ".", ".", ".", ".", ".", ".", ".", "."},
		{"6", ".", ".", ".", ".", ".", ".", ".", "."},
		{"7", ".", ".", ".", ".", ".", ".", ".", "."},
		{"8", ".", ".", ".", ".", ".", ".", ".", "."},
		{"9", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	arr1 := make([][]byte, 0)

	for _, v := range arr {
		tmp := make([]byte, 0)
		for _, v1 := range v {
			r := []rune(v1)
			rr := byte(r[0])
			tmp = append(tmp, rr)
		}
		arr1 = append(arr1, tmp)
	}

	fmt.Println(isValidSudoku(arr1))
}

func isValidSudoku(board [][]byte) bool {

	m := make(map[byte][][]int)

	for k, v := range board {
		for k1, v1 := range v {
			if v1 == 46 {
				continue
			} else {
				if tmp, ok := m[v1]; ok {
					for _, vv := range tmp {
						hang := vv[0]
						lie := vv[1]
						if k == hang || k1 == lie || (hang/3 == k/3 && lie/3 == k1/3) {
							return false
						} else {
							m[v1] = append(m[v1], []int{k, k1})
						}
					}

				} else {
					m[v1] = append(m[v1], []int{k, k1})
				}
			}
		}
	}
	return true
}
