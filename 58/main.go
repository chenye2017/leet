package main

import "fmt"

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	var tmp string
	arr := make([]string, 0)
	for _, v := range []rune(s) {
		if string(v) == " " {
			if len(tmp) > 0 {
				arr = append(arr, tmp)
				tmp = ""
			}
		} else {
			tmp += string(v)
		}
	}
	if len(tmp) > 0 {
		arr = append(arr, tmp)
	}

	return len(arr[len(arr)-1])
}
