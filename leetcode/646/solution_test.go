package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(findLongestChain([][]int{
		{-6, 9}, {1, 6}, {8, 10}, {-1, 4}, {-6, -2}, {-9, 8}, {-5, 3}, {0, 3},
	}))
	t.Log(findLongestChain([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}))
	t.Log(findLongestChain([][]int{
		{1, 2},
		{7, 8},
		{4, 5},
	}))
}
