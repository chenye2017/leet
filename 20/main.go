package main

import "fmt"

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	n := map[string]struct{}{"[": {}, "{": {}, "(": {}}
	m := map[string]string{"]": "[", "}": "{", ")": "("}

	arr := make([]string, 0)

	for _, v := range []rune(s) {
		if _, ok := n[string(v)]; ok {
			arr = append(arr, string(v))
		} else {
			if len(arr) == 0 {
				return false
			}
			tmp1 := arr[len(arr)-1]
			if tmp1 != m[string(v)] {
				return false
			}
			arr = arr[0 : len(arr)-1]
		}
	}

	if len(arr) > 0 {
		return false
	}

	return true
}
