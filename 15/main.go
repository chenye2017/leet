package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	nums = []int{0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	res := make(map[int]struct{})
	count := len(nums)
	if count < 3 {
		return [][]int{}
	}

	targetFunc := func(arr []int, target int) [][]int {
		m := make(map[int]struct{})      // 元素map
		resMap := make(map[int]struct{}) // 结果map
		res := make([][]int, 0)

		for _, v := range arr {
			tmp := target - v
			if _, ok := m[tmp]; ok {
				if _, ok := resMap[tmp]; !ok {
					res = append(res, []int{v, tmp})
					resMap[tmp] = struct{}{}
					resMap[v] = struct{}{}
				}
			} else {
				m[v] = struct{}{}
			}
		}
		return res
	}

	resArr := make([][]int, 0)
	for i := 0; i <= count-3; i++ {
		if _, ok := res[i]; ok {
			continue
		}

		if company := targetFunc(nums[i+1:], -1*nums[i]); len(company) > 0 {
			for _, v := range company {
				resArr = append(resArr, append(v, nums[i]))
			}
			res[nums[i]] = struct{}{}
		}
	}

	endArr := make([][]int, 0)
	endM := make([]map[int]struct{}, 0)

	for k := range resArr {
		sort.SliceStable(resArr[k], func(i, j int) bool {
			if resArr[k][i] > resArr[k][j] {
				return true
			} else {
				return false
			}
		})
	}

	for _, v := range resArr {
		con := false
		for _, v1 := range endM {
			_, ok := v1[v[0]]
			_, ok2 := v1[v[1]]
			if ok && ok2 {
				con = true
				break
			}
		}
		if con {
			continue
		}
		endM = append(endM, map[int]struct{}{v[0]: {}, v[1]: {}, v[2]: {}})
		endArr = append(endArr, v)
	}

	return endArr
}
