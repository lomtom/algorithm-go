package leetcode

// 执行耗时:6 ms,击败了75.69% 的Go用户
// 内存消耗:6 MB,击败了64.67% 的Go用户
func kthSmallest(root *TreeNode, k int) int {
	var stack = make([]*TreeNode, 0)
	stack = append(stack, root)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return 0
}
