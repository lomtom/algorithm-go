package leetcode

func averageOfSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return averageOfSubtree(root.Left) + averageOfSubtree(root.Right)
}
