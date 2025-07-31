package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt32
	var prev *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil && node.Val-prev.Val < ans {
			ans = node.Val - prev.Val
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
