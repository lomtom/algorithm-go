package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maximumDetonation([][]int{
		{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4},
	}))
	t.Log(maximumDetonation([][]int{
		{1, 1, 5}, {10, 10, 5},
	}))
	t.Log(maximumDetonation([][]int{
		{2, 1, 3}, {6, 1, 4},
	}))
}
