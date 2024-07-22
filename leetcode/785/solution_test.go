package leetcode

import "testing"

func Test(t *testing.T) {
	// true
	t.Log(isBipartite([][]int{
		{1}, {0}, {3}, {2},
	}))
	t.Log(isBipartite([][]int{
		{4}, {}, {4}, {4}, {0, 2, 3},
	}))
	t.Log(isBipartite([][]int{
		{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2},
	}))
	t.Log(isBipartite([][]int{
		{1, 3}, {0, 2}, {1, 3}, {0, 2},
	}))
	t.Log(isBipartite([][]int{
		{},
	}))
	t.Log(isBipartite([][]int{
		{1}, {0},
	}))
	t.Log(isBipartite([][]int{
		{1}, {0}, {3}, {2}, {5}, {4},
	}))
}
