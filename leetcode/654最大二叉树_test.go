package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	l := len(nums)
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		m[nums[i]] = i
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	root := &TreeNode{
		Val: nums[0],
	}
	var insertNode func(node *TreeNode, num int)
	insertNode = func(node *TreeNode, num int) {
		if m[node.Val] > m[num] {
			if node.Left == nil {
				node.Left = &TreeNode{
					Val: num,
				}
			} else {
				insertNode(node.Left, num)
			}
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{
					Val: num,
				}
			} else {
				insertNode(node.Right, num)
			}
		}
	}
	for i := 1; i < l; i++ {
		insertNode(root, nums[i])
	}
	return root
}

//执行用时：40 ms, 在所有 Go 提交中击败了6.08%的用户
//内存消耗：6.9 MB, 在所有 Go 提交中击败了95.77%的用户
func TestConstructMaximumBinaryTree(t *testing.T) {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(root)
}
