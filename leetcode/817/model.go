package leetcode

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func constructListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val: nums[0],
	}
	var Func func(node *ListNode, index int)
	Func = func(node *ListNode, index int) {
		if index < len(nums) {
			node.Next = &ListNode{
				Val: nums[index],
			}
			index++
			Func(node.Next, index)
		}
	}
	Func(root, 1)
	return root
}
