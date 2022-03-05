package main

import "fmt"

func main() {
	a := "aba"
	b := "cdc"
	fmt.Println(findLUSlength2(a, b))
}

func findLUSlength2(a string, b string) int {
	countA := len(a)
	countB := len(b)

	if countA == 0 {
		return len(b)
	}

	if countB == 0 {
		return len(a)
	}

	if a == b {
		return -1
	}

	var loop string
	var m = make(map[string]int)
	var mLoop string
	if countA > countB {
		loop = b
		mLoop = a
	} else {
		loop = a
		mLoop = b
	}

	if countA != countB {
		return len(mLoop)
	}

	for k, v := range mLoop {
		if k <= m[string(v)] {
			m[string(v)] = k
		}
	}

	if position, ok := m[string(loop[0])]; !ok {
		return len(mLoop)
	} else {
		if position+len(loop) > len(mLoop) {
			return len(mLoop)
		} else {
			if mLoop[position:position+len(loop)] != loop {
				return len(mLoop)
			} else {
				max := position
				end := len(mLoop) - position - len(loop)

				if max < end {
					max = end
				}
				if max < len(loop) {
					max = len(loop)
				}
				return max
			}
		}
	}
}
