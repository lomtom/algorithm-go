package leetcode

// 执行耗时:4 ms,击败了23.66% 的Go用户
// 内存消耗:2.6 MB,击败了20.79% 的Go用户
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(p, q *TreeNode) bool {
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	return p.Val == q.Val && isMirror(p.Right, q.Left) && isMirror(p.Left, q.Right)
}
