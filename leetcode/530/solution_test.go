package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(getMinimumDifference(construct([]int{5, 4, 7})))
	t.Log(getMinimumDifference(construct([]int{4, 2, 6, 1, 3})))
}
