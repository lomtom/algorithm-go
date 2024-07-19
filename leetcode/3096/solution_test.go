package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(minimumLevels([]int{1, 1}))
	t.Log(minimumLevels([]int{1, 1, 1, 1}))
	t.Log(minimumLevels([]int{1, 0, 1, 0}))
	t.Log(minimumLevels([]int{0, 0}))
}
