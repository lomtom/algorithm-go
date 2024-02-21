package leetcode

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
