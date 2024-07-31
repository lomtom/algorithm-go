package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(minRectanglesToCoverPoints([][]int{
		{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6},
	}, 2))
	t.Log(minRectanglesToCoverPoints([][]int{
		{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6},
	}, 1))
	t.Log(minRectanglesToCoverPoints([][]int{
		{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6},
	}, 2))
	t.Log(minRectanglesToCoverPoints([][]int{
		{2, 3}, {1, 2},
	}, 0))
}
