package leetcode

import "sort"

// 执行耗时:40 ms,击败了10.00% 的Go用户
// 内存消耗:10.4 MB,击败了5.88% 的Go用户
func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	if special[0] != bottom {
		special = append([]int{bottom - 1}, special...)
	}
	if special[len(special)-1] != top {
		special = append(special, top+1)
	}
	var result int
	for i := 1; i < len(special); i++ {
		result = max(result, special[i]-special[i-1]-1)
	}
	return result
}
