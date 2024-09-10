package leetcode

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

func antiConstructListNode(root *ListNode) []int {
	var nums []int
	for root != nil {
		nums = append(nums, root.Val)
		root = root.Next
	}
	return nums
}
