package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了37.31% 的Go用户
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
