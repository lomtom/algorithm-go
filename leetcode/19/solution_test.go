package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(antiConstructListNode(removeNthFromEnd(constructListNode([]int{1}), 1)))
	t.Log(antiConstructListNode(removeNthFromEnd(constructListNode([]int{1, 2, 3, 4, 5}), 2)))
	t.Log(antiConstructListNode(removeNthFromEnd(constructListNode([]int{1, 2}), 1)))
	t.Log(antiConstructListNode(removeNthFromEnd(constructListNode([]int{1, 2}), 2)))
}
