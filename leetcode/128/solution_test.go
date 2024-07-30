package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(longestConsecutive([]int{1, 5}))
	t.Log(longestConsecutive([]int{1, 3, 5, 2, 4}))
	t.Log(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	t.Log(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
