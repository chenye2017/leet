package main

import "fmt"

func main() {
	r := "bg"
	c := "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"

	fmt.Println(canConstruct(r, c))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	rm := make(map[rune]int)

	for _, v := range ransomNote {
		rm[v]++
	}

	mm := make(map[rune]int)

	for _, v := range magazine {
		mm[v]++
	}

	for k, v := range rm {
		if tmp, ok := mm[k]; !ok {
			return false
		} else {
			if v > tmp {
				return false
			}
		}
	}
	return true

}
