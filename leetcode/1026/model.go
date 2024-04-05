package leetcode

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(nums []any) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var Func func(index int) *TreeNode
	Func = func(index int) *TreeNode {
		if index < len(nums) {
			val := nums[index]
			switch v := val.(type) {
			case int:
				// 当值为int类型时，创建新的树节点
				node := &TreeNode{
					Val:   v,
					Left:  Func(2*index + 1),
					Right: Func(2*index + 2),
				}
				return node
			case nil:
				// 当值为nil时，不创建子节点
				break
			}
		}
		return nil
	}
	return Func(0)
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
