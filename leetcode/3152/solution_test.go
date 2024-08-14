package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(isArraySpecial([]int{3, 4, 1, 2, 6}, [][]int{
		{0, 4},
	}))
	t.Log(isArraySpecial([]int{4, 3, 1, 6}, [][]int{
		{0, 2},
		{2, 3},
	}))
}
