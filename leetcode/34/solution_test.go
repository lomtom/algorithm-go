package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(searchRange([]int{1}, 1))
	t.Log(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	t.Log(searchRange([]int{3, 4, 5}, 3))
	t.Log(searchRange([]int{3, 4, 5}, 8))
	t.Log(searchRange([]int{3, 4, 5}, 1))
	t.Log(searchRange([]int{}, 0))
}
