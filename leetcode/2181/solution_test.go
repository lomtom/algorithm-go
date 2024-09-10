package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(antiConstructListNode(mergeNodes(constructListNode([]int{0, 3, 1, 0, 4, 5, 2, 0}))))
	t.Log(antiConstructListNode(mergeNodes(constructListNode([]int{0, 1, 0, 3, 0, 2, 2, 0}))))
}
