package main

import "fmt"

func main() {
	a := "aba"
	b := "cdc"
	fmt.Println(findLUSlength1(a, b))
}

func findLUSlength1(a string, b string) int {
	countA := len(a)
	countB := len(b)

	if countA == 0 {
		return len(b)
	}

	if countB == 0 {
		return len(a)
	}

	if a == b {
		return -1
	}

	var loop string
	var m = make(map[string][]int)
	var mLoop string
	if countA > countB {
		loop = b
		mLoop = a
	} else {
		loop = a
		mLoop = b
	}

	for k, v := range mLoop {
		m[string(v)] = append(m[string(v)], k)
	}

	nowStart := 0
	res := make([]string, 0)
	str := ""
	for _, v := range loop {
		if arr, ok := m[string(v)]; !ok {
			continue
		} else {
			flag := false
			for k, v1 := range arr {
				if nowStart <= v1 {
					nowStart = v1
					str += string(v)
					flag = true
					m[string(v)] = append(m[string(v)][0:k], m[string(v)][k+1:]...)
					break
				}
			}

			if !flag {
				res = append(res, str)
				str = ""
				nowStart = 0
			}
		}
	}

	if len(str) > 0 {
		res = append(res, str)
	}

	if len(res) == 0 {
		return len(mLoop)
	} else {
		max := len(res[0])
		for i := 1; i < len(res); i++ {
			if len(res[i]) > max {
				max = len(res[i])
			}
		}
		return len(mLoop) - max
	}
}
