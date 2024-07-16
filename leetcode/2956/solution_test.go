package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(findIntersectionValues([]int{4, 3, 2, 3, 1}, []int{2, 2, 5, 2, 3, 6}))
	t.Log(findIntersectionValues([]int{3, 4, 2, 3}, []int{1, 5}))
}
