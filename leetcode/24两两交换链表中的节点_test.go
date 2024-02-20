package leetcode

import "testing"

// [1,2,3,4] -> [2,1,4,3]
// 执行耗时:1 ms,击败了10.86% 的Go用户
// 内存消耗:2 MB,击败了76.34% 的Go用户
func swapPairs(head *ListNode) *ListNode {
	root := &ListNode{}
	root.Next = head
	pre := root
	cur := root.Next
	for cur != nil && cur.Next != nil {
		temp := cur.Next.Next
		pre.Next = cur.Next
		pre.Next.Next = cur
		cur.Next = temp
		pre = cur
		cur = temp
	}
	return root.Next
}

func TestSwapPairs(t *testing.T) {
	t.Log(antiConstructListNode(swapPairs(constructListNode([]int{1, 2, 3, 4, 5}))))
	t.Log(antiConstructListNode(swapPairs(constructListNode([]int{}))))
	t.Log(antiConstructListNode(swapPairs(constructListNode([]int{1}))))
}
