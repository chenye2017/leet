package main

import "fmt"

func main() {
	s := "badc"
	t := "baba"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]byte)
	inUse := make(map[byte]struct{})

	for k, v := range s {
		v1 := byte(v)
		v2 := t[k]
		if _, ok := m[v1]; !ok {
			if _, ok := inUse[v2]; ok {
				return false
			}
			m[v1] = v2
			inUse[v2] = struct{}{}
		} else {
			if m[v1] != v2 {
				return false
			}
		}
	}
	return true
}
