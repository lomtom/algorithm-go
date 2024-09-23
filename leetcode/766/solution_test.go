package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
	t.Log(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}
