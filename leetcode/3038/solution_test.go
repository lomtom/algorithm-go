package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maxOperations([]int{2, 2, 3, 2, 4, 2, 3, 3, 1, 3}))
	t.Log(maxOperations([]int{1, 2}))
	t.Log(maxOperations([]int{3, 2, 1, 4, 5}))
	t.Log(maxOperations([]int{3, 2, 6, 1, 4}))
}
