package leetcode

// 执行耗时:4 ms,击败了92.86% 的Go用户
// 内存消耗:6.1 MB,击败了83.33% 的Go用户
func findTilt(root *TreeNode) int {
	var sum int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		var leftValue = dfs(node.Left)
		var rightValue = dfs(node.Right)
		if leftValue < rightValue {
			sum += rightValue - leftValue
		} else {
			sum += leftValue - rightValue
		}
		return node.Val + leftValue + rightValue
	}
	dfs(root)
	return sum
}
