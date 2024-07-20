package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(minimumMoves([][]int{{1, 3, 0}, {1, 0, 0}, {1, 0, 3}}))
	t.Log(minimumMoves([][]int{{1, 1, 0}, {1, 1, 1}, {1, 2, 1}}))
}
