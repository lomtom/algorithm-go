package _19

import (
	"fmt"
	"testing"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	queue := make([]*TreeNode, 0)
	tempQueue := []*TreeNode{root}
	for len(tempQueue) > 0 {
		node := tempQueue[0]
		tempQueue = tempQueue[1:]
		queue = append(queue, node)
		if node.Left != nil {
			tempQueue = append(tempQueue, node.Left)
		}
		if node.Right != nil {
			tempQueue = append(tempQueue, node.Right)
		}
	}
	return CBTInserter{
		root,
		queue,
	}
}

func (this *CBTInserter) Insert(val int) int {
	node := this.queue[0]
	if node.Left != nil && node.Right != nil {
		this.queue = this.queue[1:]
		return this.Insert(val)
	}
	child := &TreeNode{
		Val: val,
	}
	if node.Left == nil {
		node.Left = child
	} else if node.Right == nil {
		node.Right = child
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, child)
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

//执行用时：4 ms, 在所有 Go 提交中击败了94.59%的用户
//内存消耗：6.6 MB, 在所有 Go 提交中击败了35.14%的用户
func TestCBTInserter(t *testing.T) {
	obj := Constructor(&TreeNode{Val: 1})
	param_1 := obj.Insert(2)
	param_1 = obj.Insert(3)
	param_1 = obj.Insert(4)
	param_2 := obj.Get_root()
	fmt.Println(param_1)
	fmt.Println(param_2)
}
