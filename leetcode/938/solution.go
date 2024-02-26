package leetcode

// 深度优先搜索（递归）
// 执行耗时:68 ms,击败了76.92% 的Go用户
// 内存消耗:9.7 MB,击败了61.54% 的Go用户
func rangeSumBST(root *TreeNode, low int, high int) (result int) {
	if root == nil {
		return 0
	}
	if root.Val <= high && root.Val >= low {
		result += root.Val
	}
	return rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high) + result
}

// 深度优先搜索（迭代）
// 执行耗时:72 ms,击败了48.72% 的Go用户
// 内存消耗:10 MB,击败了17.95% 的Go用户
func rangeSumBST1(root *TreeNode, low int, high int) (result int) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			if root.Val <= high && root.Val >= low {
				result += root.Val
			}
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return
}
