package leetcode

import "testing"

func Test(t *testing.T) {
	// 9
	t.Log(maxScore([][]int{
		{5, 4, 8, 6, 2},
		{1, 5, 8, 6, 10},
		{8, 6, 9, 2, 10},
		{3, 7, 6, 10, 6},
	}))
	// 9
	t.Log(maxScore([][]int{
		{9, 5, 7, 3},
		{8, 9, 6, 1},
		{6, 7, 14, 3},
		{2, 5, 3, 1},
	}))
	t.Log(maxScore([][]int{
		{4, 3, 2},
		{3, 2, 1},
	}))
}
