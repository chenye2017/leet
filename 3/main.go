package main

import "fmt"

func main() {
	s := "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}

/**
连续子串，想到的就是滑动窗口
每次替换的时候记着max长度
关于移动前置窗口索引，老师的方法是一个个移动
我是记录出位置，整体把前置index往前移动
*/
func lengthOfLongestSubstring(s string) int {
	count := len(s)

	if count == 0 {
		return 0
	}

	start := 0
	end := 0
	m := make(map[byte]int)
	max := 0
	now := 0
	for {
		if tmp, ok := m[s[end]]; !ok {
			now++
		} else {
			if max < now {
				max = now
			}
			start = tmp + 1
			for k, v := range m {
				if v < start {
					delete(m, k)
				}
			}

			now = end - start + 1
		}
		m[s[end]] = end
		end++
		if end >= count {
			break
		}
	}

	if now > max {
		max = now
	}

	return max
}
