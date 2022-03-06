package main

func main() {

}

func minWindow(s string, t string) string {
	cs := len(s)
	ct := len(t)

	if cs < ct {
		return ""
	}

	mt := make(map[rune]int)

	for _, v := range s {
		if _, ok := mt[v]; ok {
			mt[v] += 1
		} else {
			mt[v] = 1
		}
	}

	start := 0
	for k, v := range s {
		if _, ok := mt[v]; ok {
			start = k
			break
		}
	}

	end := start + ct - 1

	if end > cs-1 {
		return ""
	}

	tmp := s[start : end+1]

	for _, v := range tmp {

	}

	for {

	}

}
