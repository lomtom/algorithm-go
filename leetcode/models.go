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

func antiConstructListNode(root *ListNode) []int {
	var nums []int
	for root != nil {
		nums = append(nums, root.Val)
		root = root.Next
	}
	return nums
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(nums []int) *TreeNode {
	root := &TreeNode{
		Val: nums[0],
	}
	var Func func(node *TreeNode, index int)
	Func = func(node *TreeNode, index int) {
		if 2*index+1 < len(nums) {
			node.Left = &TreeNode{
				Val: nums[2*index+1],
			}
			Func(node.Left, 2*index+1)
		}
		if 2*index+2 < len(nums) {
			node.Right = &TreeNode{
				Val: nums[2*index+2],
			}
			Func(node.Right, 2*index+2)
		}
	}
	Func(root, 0)
	return root
}
