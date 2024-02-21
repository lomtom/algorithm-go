package leetcode

import "testing"

func TestSwapPairs(t *testing.T) {
	t.Log(antiConstructListNode(swapPairs(constructListNode([]int{1, 2, 3, 4, 5}))))
	t.Log(antiConstructListNode(swapPairs(constructListNode([]int{}))))
	t.Log(antiConstructListNode(swapPairs(constructListNode([]int{1}))))
}
