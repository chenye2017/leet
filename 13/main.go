package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	mi := map[string]int{
		"V": 4,
		"X": 9,
	}
	mx := map[string]int{
		"L": 40,
		"C": 90,
	}
	mc := map[string]int{
		"D": 400,
		"M": 900,
	}

	mm := map[string]map[string]int{
		"I": mi,
		"X": mx,
		"C": mc,
	}

	count := len([]rune(s))

	total := 0

	next := false
	for k, v := range []rune(s) {
		if next {
			next = false
			continue
		}

		currentP := string(v)

		if tmp, ok := mm[currentP]; ok {
			if k == count-1 {
				total += m[currentP]
			} else {
				nextP := string([]rune(s)[k+1])
				if tmp1, ok := tmp[nextP]; ok {
					next = true
					total += tmp1
				} else {
					total += m[currentP]
				}
			}

		} else {
			total += m[currentP]
		}
	}
	return total
}
