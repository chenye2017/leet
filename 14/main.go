package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	minIndex := 0
	minCount := len(strs[0])

	for k, v := range strs {
		if k == 0 {
			continue
		} else {
			tmp := len(v)
			if tmp < minCount {
				minIndex = k
				minCount = tmp
			}
		}
	}

	mm := make([]map[byte]int)
	for k, v := range strs {
		if k == minIndex {
			continue
		} else {
			tmp := make(map[byte]int)
			for k, v1 := range []byte(v) {
				tmp[v1]
			}
		}
	}

}
