package main

import (
	"fmt"
	"sort"
)

func main() {
	index := nextGreatestLetter([]byte{byte('c'), byte('f'), byte('j')}, byte('j'))
	fmt.Println(string(index))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	count := len(letters)
	index := sort.Search(count, func(i int) bool {
		return letters[i] > target
	})

	if index < count {
		return letters[index]
	} else {
		if letters[count-1] > target {
			return letters[count-1]
		} else {
			return letters[0]
		}
	}
}
