package leetcode

// 执行耗时:55 ms,击败了93.83% 的Go用户
// 内存消耗：7.3 MB,击败了93.28% 的Go用户
func longestUnivaluePath(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		if node.Left != nil && node.Left.Val == node.Val {
			l = l + 1
		} else {
			l = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
			r = r + 1
		} else {
			r = 0
		}
		ans = max(ans, l+r)
		return max(l, r)
	}
	dfs(root)
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
