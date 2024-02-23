package leetcode

import "testing"

func TestSwapPairs(t *testing.T) {
	t.Log(antiConstructListNode(mergeTwoLists(constructListNode([]int{1, 2, 4}), constructListNode([]int{1, 3, 4}))))
}
