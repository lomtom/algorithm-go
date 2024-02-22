package leetcode

import "testing"

func TestSwapPairs(t *testing.T) {
	t.Log(antiConstructListNode(reverseList(constructListNode([]int{1, 2, 3, 4, 5}))))
	t.Log(antiConstructListNode(reverseList(constructListNode([]int{}))))
	t.Log(antiConstructListNode(reverseList(constructListNode([]int{1}))))
}
