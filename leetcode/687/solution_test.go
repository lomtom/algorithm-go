package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(longestUnivaluePath(construct([]int{5, 4, 5, 1, 1, 5})))
	t.Log(longestUnivaluePath(construct([]int{1, 4, 5, 4, 4, 5})))
}
