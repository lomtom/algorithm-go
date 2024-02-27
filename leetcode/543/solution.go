package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:4.4 MB,击败了80.62% 的Go用户
func diameterOfBinaryTree(root *TreeNode) int {
	var ans = 1
	getMaxDiameter(root, &ans)
	return ans - 1
}

func getMaxDiameter(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := getMaxDiameter(root.Left, ans)
	right := getMaxDiameter(root.Right, ans)
	*ans = max(*ans, left+right+1)
	return max(left, right) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
