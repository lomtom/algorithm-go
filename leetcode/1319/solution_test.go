package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	t.Log(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
	t.Log(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}))
	t.Log(makeConnected(5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}))
}
