package leetcode

import (
	"fmt"
	"strings"
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

// printHorizontalTree 打印横向的二叉树图形
func printHorizontalTree(root *TreeNode, level int, isLeft bool, prefix string) {
	if root == nil {
		return
	}

	if level > 0 {
		fmt.Print(strings.Repeat("    ", level-1))
		if isLeft {
			fmt.Print("└───")
		} else {
			fmt.Print("├───")
		}
	}

	fmt.Println(root.Val)

	if root.Left != nil || root.Right != nil {
		printHorizontalTree(root.Left, level+1, true, prefix+"|   ")
		printHorizontalTree(root.Right, level+1, false, prefix+"|   ")
	}
}
