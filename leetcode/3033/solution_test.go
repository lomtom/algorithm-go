package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(modifiedMatrix([][]int{
		{1, 2, -1}, {4, -1, 6}, {7, 8, 9},
	}))
	t.Log(modifiedMatrix([][]int{
		{3, -1}, {5, 2},
	}))
}
