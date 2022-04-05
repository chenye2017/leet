package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	matrix = [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	defer func() {
		fmt.Println(matrix)
	}()
	mHang := len(matrix)

	if mHang == 0 {
		return
	}
	mLie := len(matrix[0])
	hang := make([]int, 0)
	lie := make([]int, 0)
	for k, v := range matrix {
		for k1, v1 := range v {
			if v1 == 0 {
				hang = append(hang, k)
				lie = append(lie, k1)
			}
		}
	}

	for _, v := range hang {
		for i := 0; i < mLie; i++ {
			matrix[v][i] = 0
		}
	}
	for _, v := range lie {
		for i := 0; i < mHang; i++ {
			matrix[i][v] = 0
		}
	}
}
