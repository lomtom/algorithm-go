package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(firstMissingPositive([]int{1, 2, 0}))
	t.Log(firstMissingPositive([]int{3, 4, -1, 1}))
	t.Log(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
