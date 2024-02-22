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

func antiConstruct(root *TreeNode) []int {
	var res []int
	// 层序遍历保存结果
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node != nil {
			res = append(res, node.Val)
			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
	}
	return res
}
