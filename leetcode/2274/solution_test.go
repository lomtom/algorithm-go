package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maxConsecutive(2, 9, []int{4, 6}))
	t.Log(maxConsecutive(6, 8, []int{7, 6, 8}))
	t.Log(maxConsecutive(1, 2, []int{1, 2}))
	t.Log(maxConsecutive(1, 1, []int{1}))
}
