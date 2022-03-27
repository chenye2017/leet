package main

func main() {

}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	count := m * n
	if count != r*c {
		return mat
	}

	res := make([]int, 0, count)
	for _, v := range mat {
		for _, v1 := range v {
			res = append(res, v1)
		}
	}

	re := make([][]int, 0)
	for i := 0; i < r; i++ {
		re = append(re, res[i*c:i*c+c])
	}
	return re
}
