package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 15, 11, 2}
	b := []int{23, 127, 235, 19, 8}

	fmt.Println(smallestDifference(a, b))
}

func smallestDifference(a []int, b []int) int {
	sort.SliceStable(a, func(i, j int) bool {
		if a[i] < a[j] {
			return true
		} else {
			return false
		}
	})

	sort.SliceStable(b, func(i, j int) bool {
		if b[i] < b[j] {
			return true
		} else {
			return false
		}
	})

	aIndex := 0
	acount := len(a)
	bIndex := 0
	bcount := len(b)

	gap := -1

	var flag bool
	for {
		tmp := a[aIndex] - b[bIndex]

		if tmp >= 0 {
			bIndex++
			if bIndex >= bcount {
				flag = true
			}
		} else {
			aIndex++
			tmp = tmp * -1
			if aIndex >= acount {
				flag = true
			}
		}

		if gap == -1 || gap > tmp {
			gap = tmp
		}

		if flag {
			break
		}
	}

	return gap
}
