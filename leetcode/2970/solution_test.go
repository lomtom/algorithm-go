package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(incremovableSubarrayCount([]int{1, 2, 3, 4, 5}))
	t.Log(incremovableSubarrayCount([]int{6, 5, 7, 8}))
	t.Log(incremovableSubarrayCount([]int{8, 7, 6, 6}))
}
