package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(relocateMarbles([]int{1, 6, 7, 8}, []int{1, 7, 2}, []int{2, 9, 5}))
	t.Log(relocateMarbles([]int{1, 1, 3, 3}, []int{1, 3}, []int{2, 2}))
}
