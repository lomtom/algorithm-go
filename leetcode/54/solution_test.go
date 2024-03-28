package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(spiralOrder([][]int{
		{3},
		{2},
	}))
	t.Log(spiralOrder([][]int{
		{3, 2},
	}))
	t.Log(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	t.Log(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
	t.Log(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}))
}
