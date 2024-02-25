package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(kthLargestLevelSum(construct([]int{5, 8, 9, 2, 1, 3, 7}), 4))
	t.Log(kthLargestLevelSum(construct([]int{5, 8, 9, 2, 1, 3, 7, 4, 6}), 2))
}
