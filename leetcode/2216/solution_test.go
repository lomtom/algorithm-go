package leetcode

import "testing"

func Test(t *testing.T) {
	// 4
	t.Log(minDeletion([]int{0, 1, 5, 4, 2, 4, 7, 2, 3, 0, 3, 0, 0, 9, 7, 5, 9, 4, 3, 9, 9, 2, 1, 6, 3, 1, 0, 7, 6, 6, 6, 0, 1, 7, 1, 9, 4, 9, 3, 3, 4, 5, 0, 3, 8, 7, 1, 8, 4, 5}))
	t.Log(minDeletion([]int{1, 1, 2, 3, 5}))
	t.Log(minDeletion([]int{1, 1, 2, 2, 3, 3}))
}
