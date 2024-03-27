package leetcode

import (
	"sort"
)

const mod = 1e9 + 7

// 执行耗时:127 ms,击败了75.00% 的Go用户
// 内存消耗:15.9 MB,击败了91.67% 的Go用户
func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})
	var count = 1
	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] > ranges[i-1][1] {
			count++
		} else if ranges[i][1] < ranges[i-1][1] {
			ranges[i][1] = ranges[i-1][1]
		}
	}
	return pow(2, count)
}

func pow(x, num int) int {
	var result = 1
	for num > 0 {
		if num%2 == 1 {
			result *= x
			result %= mod
		}
		x *= x
		x %= mod
		num = num >> 1
	}
	return result
}
