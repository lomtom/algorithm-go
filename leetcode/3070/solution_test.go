package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(countSubmatrices([][]int{
		{7, 6, 3}, {6, 6, 1},
	}, 18))
}
