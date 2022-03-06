package main

import "fmt"

func main() {
	s := "rat"
	t := "car"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	s1 := []byte(s)
	s2 := []byte(t)

	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for _, v := range s1 {
		m1[v]++
	}
	for _, v := range s2 {
		m2[v]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if tmp, ok := m2[k]; !ok {
			return false
		} else {
			if tmp != v {
				return false
			}
		}
	}
	return true
}
