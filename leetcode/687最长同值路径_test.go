package leetcode

//执行用时：56 ms, 在所有 Go 提交中击败了64.93%的用户
//内存消耗：7.3 MB, 在所有 Go 提交中击败了93.28%的用户
func longestUnivaluePath(root *TreeNode) (ans int) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.Left), dfs(node.Right)
		if node.Left != nil && node.Left.Val == node.Val {
			left = left + 1
		} else {
			left = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
			right = right + 1
		} else {
			right = 0
		}
		ans = max(ans, left+right)
		return max(left, right)
	}
	dfs(root)
	return
}
