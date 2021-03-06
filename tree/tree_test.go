package tree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestPreTraversal(t *testing.T) {
	root := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4, nil, nil,
			},
			&TreeNode{
				5, nil, nil,
			},
		},
		&TreeNode{
			3,
			&TreeNode{
				6, nil, nil,
			},
			&TreeNode{
				7, nil, nil,
			},
		},
	}
	preTraversal(root)
	fmt.Println()
	preTraversal1(root)
	fmt.Println()
	miTraversal(root)
	fmt.Println()
	miTraversal1(root)
	fmt.Println()
	nextTraversal(root)
	fmt.Println()
	nextTraversal1(root)
	fmt.Println()
	levelTraversal(root)
	fmt.Println()
	fmt.Println(dfs(root))
}

//1	 2	4	5	3	6	7
func preTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, "\t")
	preTraversal(root.Left)
	preTraversal(root.Right)
}

func preTraversal1(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			fmt.Print(root.Val, "\t")
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
}

func miTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	miTraversal(root.Left)
	fmt.Print(root.Val, "\t")
	miTraversal(root.Right)
}

func miTraversal1(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(root.Val, "\t")
		root = root.Right
	}
}

func nextTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	nextTraversal(root.Left)
	nextTraversal(root.Right)
	fmt.Print(root.Val, "\t")
}

func nextTraversal1(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			fmt.Print(node.Val, "\t")
			// ???????????????????????????????????????
			lastVisit = node
		} else {
			root = node.Right
		}
	}
}

func levelTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		fmt.Print(root.Val, "\t")
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if len(stack) == 0 {
			root = nil
		} else {
			root = stack[0]
			stack = stack[1:]
		}
	}
}

// ??????
func dfs(root *TreeNode) (result []int) {
	result = make([]int, 0)
	if root == nil {
		return
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	result = append(result, left...)
	result = append(result, right...)
	result = append(result, root.Val)
	return
}

// maxDepth https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func TestMaxDepth(t *testing.T) {

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
