package leetcode

import "sort"

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.27 MB,击败了32.00% 的Go用户
func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	for start := len(citations); start > 0; start-- {
		if citations[start-1] >= start {
			return start
		}
	}
	return 0
}
