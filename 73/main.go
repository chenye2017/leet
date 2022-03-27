package main

func main() {

}

func setZeroes(matrix [][]int) {
	n := len(matrix[0])

	hang := make(map[int]struct{})
	lie := make(map[int]struct{})

	for k, v := range matrix {
		for k1, v1 := range v {
			if v1 == 0 {
				hang[k] = struct{}{}
				lie[k1] = struct{}{}
			}
		}
	}

	for k := range hang {
		matrix[k] = make([]int, n)
	}
	for k := range lie {
		for k1 := range matrix {
			matrix[k1][k] = 0
		}
	}
	return
}
