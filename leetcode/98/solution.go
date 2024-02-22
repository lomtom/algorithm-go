package leetcode

// 执行耗时:3 ms,击败了86.13% 的Go用户
// 内存消耗:6.1 MB,击败了7.58% 的Go用户
func isValidBST(root *TreeNode) bool {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}
	return true
}

// 执行耗时:3 ms,击败了86.13% 的Go用户
// 内存消耗:5.2 MB,击败了68.72% 的Go用户
func isValidBST1(root *TreeNode) bool {
	var helper func(node *TreeNode, min *TreeNode, max *TreeNode) bool
	helper = func(node *TreeNode, min *TreeNode, max *TreeNode) bool {
		if node == nil {
			return true
		}
		if min != nil && node.Val <= min.Val {
			return false
		}
		if max != nil && node.Val >= max.Val {
			return false
		}
		return helper(node.Left, min, node) && helper(node.Right, node, max)
	}
	return helper(root, nil, nil)
}
