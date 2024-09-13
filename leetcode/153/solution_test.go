package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(findMin([]int{3, 4, 5, 6, 7, 0, 1, 2}))
	t.Log(findMin([]int{5, 6, 7, 0, 1, 2, 3, 4}))
	t.Log(findMin([]int{2, 3, 4, 5, 6, 7, 0, 1}))
	t.Log(findMin([]int{6, 7, 0, 1, 2, 3, 4, 5}))
	t.Log(findMin([]int{0, 1, 2, 3, 4, 5, 6, 7}))
}
