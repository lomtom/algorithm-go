package leetcode

import (
	"sort"
	"testing"
)

// 执行耗时:21 ms,击败了17.08% 的Go用户
// 内存消耗:6.1 MB,击败了73.13% 的Go用户
func merge1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] > intervals[j][0] {
			return false
		}
		if intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1] {
			return false
		}
		return true
	})
	result := make([][]int, 0)
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		left := result[len(result)-1]
		right := intervals[i]
		if left[1] < right[0] {
			result = append(result, right)
		} else if left[1] < right[1] {
			result[len(result)-1][1] = right[1]
		}
	}
	return result
}

func TestMerge1(t *testing.T) {
	t.Log(merge1([][]int{
		{1, 6}, {1, 5}, {0, 1}, {0, 6}, {0, 3}, {6, 9},
	}))
	t.Log(merge1([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
	t.Log(merge1([][]int{
		{1, 4}, {4, 5},
	}))
}
