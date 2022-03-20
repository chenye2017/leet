package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}
	fmt.Println(longestMountain(arr))
}

func longestMountain(arr []int) int {
	max := 0
	count := len(arr)
	if count < 3 {
		return max
	}

	up := 0
	down := 0
	for i := 1; i < count; i++ {
		tmp := arr[i] - arr[i-1]
		if tmp == 0 {
			if up > 0 && down > 0 && up+down+1 > max {
				max = up + down + 1
			}
			up = 0
			down = 0
		}
		if tmp > 0 {
			if down > 0 {
				if up > 0 && down > 0 && up+down+1 > max {
					max = up + down + 1
				}
				up = 1
				down = 0
			} else {
				up++
			}
		}
		if tmp < 0 {
			if up == 0 {
				continue
			} else {
				down++
			}
		}
	}

	if up > 0 && down > 0 && (up+down+1) > max {
		max = up + down + 1
	}

	return max
}
