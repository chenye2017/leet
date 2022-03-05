package main

func main() {

}

func findLUSlength(a string, b string) int {
	countA := len(a)
	countB := len(b)

	max := countA
	if countB > countA {
		max = countB
	}

	if countA != countB {
		return max
	} else {
		if a == b {
			return -1
		} else {
			return max
		}
	}

}
