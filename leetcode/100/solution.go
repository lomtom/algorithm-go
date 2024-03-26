package leetcode

// 执行耗时:1 ms,击败了28.52% 的Go用户
// 内存消耗:2.2 MB,击败了15.48% 的Go用户
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
