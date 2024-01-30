package leetcode

// 执行耗时:3 ms,击败了88.42% 的Go用户
// 内存消耗:4.9 MB,击败了73.17% 的Go用户
func robIII(root *TreeNode) int {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var Func func(node *TreeNode) (int, int)
	Func = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		// l1: 偷取左子树的最大值 l2: 不偷取左子树的最大值
		l1, l2 := Func(node.Left)
		// r1: 偷取右子树的最大值 r2: 不偷取右子树的最大值
		r1, r2 := Func(node.Right)
		// 偷取当前节点的最大值 = 当前节点的值 + 不偷取左子树的最大值 + 不偷取右子树的最大值
		// 不偷取当前节点的最大值 = max(偷取左子树的最大值, 不偷取左子树的最大值) + max(偷取右子树的最大值, 不偷取右子树的最大值)
		return node.Val + l2 + r2, max(l1, l2) + max(r1, r2)
	}
	return max(Func(root))
}
