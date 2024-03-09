package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(minimumOperationsToWriteY([][]int{
		{1, 2, 2}, {1, 1, 0}, {0, 1, 0},
	}))
	t.Log(minimumOperationsToWriteY([][]int{
		{0, 1, 0, 1, 0}, {2, 1, 0, 1, 2}, {2, 2, 2, 0, 1}, {2, 2, 2, 2, 2}, {2, 1, 2, 2, 2},
	}))
}
