package main

import (
	"fmt"
	"strconv"
)

func main() {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(ops))
}

func calPoints(ops []string) int {
	res := make([]int, 0)
	for _, v := range ops {
		if v == "+" {
			count := len(res)
			res = append(res, res[count-2]+res[count-1])
			continue
		}
		if v == "D" {
			count := len(res)
			res = append(res, res[count-1]*2)
			continue
		}
		if v == "C" {
			count := len(res)
			res = res[0 : count-1]
			continue
		}
		tmp, _ := strconv.Atoi(v)
		res = append(res, tmp)
	}

	sum := 0

	for _, v := range res {
		sum += v
	}
	return sum
}
