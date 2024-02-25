package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(lowestCommonAncestor(construct([]int{1, 2}), &TreeNode{Val: 1}, &TreeNode{Val: 2}))
}
