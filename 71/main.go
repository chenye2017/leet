package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/a/./b/../../c/"
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	type stack struct {
		values []string
	}
	push := func(s *stack, v string) {
		s.values = append(s.values, v)
	}
	isEmpty := func(s *stack) bool {
		if len(s.values) == 0 {
			return true
		} else {
			return false
		}
	}

	pop := func(s *stack) string {
		tmp := s.values[len(s.values)-1]
		s.values = s.values[0 : len(s.values)-1]
		return tmp
	}

	arr := strings.Split(path, "/")

	s := new(stack)
	for _, v := range arr {
		v = strings.Trim(v, "")
		if len(v) == 0 || v == "." {
			continue
		}

		if v == ".." {
			if len(s.values) == 0 {
				continue
			} else {
				pop(s)
			}
		} else {
			push(s, v)
		}
	}
	res := make([]string, len(s.values))

	start := len(s.values) - 1
	for {
		if !isEmpty(s) {
			res[start] = pop(s)
			start--
		} else {
			break
		}
	}

	return "/" + strings.Join(res, "/")
}
