package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(numberOfRightTriangles([][]int{
		{0, 1, 0}, {0, 1, 1}, {0, 1, 0},
	}))
	t.Log(numberOfRightTriangles([][]int{
		{1, 0, 0, 0}, {0, 1, 0, 1}, {1, 0, 0, 0},
	}))
	t.Log(numberOfRightTriangles([][]int{
		{1, 0, 1}, {1, 0, 0}, {1, 0, 0},
	}))
}
