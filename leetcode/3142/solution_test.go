package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(satisfiesConditions([][]int{
		{1, 0, 2},
		{1, 0, 2},
	}))
	t.Log(satisfiesConditions([][]int{
		{1, 1, 1},
		{0, 0, 0},
	}))
	t.Log(satisfiesConditions([][]int{
		{1},
		{2},
		{3},
	}))
}
