package leetcode

import (
	"fmt"
)

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
	var Func func(node *TreeNode)
	Func = func(node *TreeNode) {
		if node != nil {
			res = append(res, node.Val)
			Func(node.Left)
			Func(node.Right)
		}
	}
	Func(root)
	return res
}

// printTree 打印二叉树图形
func printTree(root *TreeNode, level int, prefix string) {
	if root == nil {
		return
	}

	if level == 0 {
		fmt.Printf("%s%d\n", prefix, root.Val)
	} else {
		fmt.Printf("%s├─── %d\n", prefix, root.Val)
	}

	if root.Left != nil || root.Right != nil {
		printTree(root.Left, level+1, prefix+"|    ")
		printTree(root.Right, level+1, prefix+"|    ")
	}
}
