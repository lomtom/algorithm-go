package leetcode

import "sort"

// 执行耗时:132 ms,击败了64.71% 的Go用户
// 内存消耗:14.9 MB,击败了17.65% 的Go用户
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	for index := range moveFrom {
		m[moveFrom[index]] = false
		m[moveTo[index]] = true
	}
	var result []int
	for k, v := range m {
		if v {
			result = append(result, k)
		}
	}
	sort.Ints(result)
	return result
}
