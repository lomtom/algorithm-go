package tree

import "testing"

// 110. 平衡二叉树 https://leetcode-cn.com/problems/balanced-binary-tree/
func TestIsBalanced(t *testing.T) {

}

//分治法:
//左边平衡 && 右边平衡 && 左右两边高度 <= 1，
//
//因为需要返回是否平衡及高度，要么返回两个数据，要么合并两个数据，
//
//所以用-1 表示不平衡，>0 表示树高度（二义性：一个变量有两种含义）。
func isBalanced(root *TreeNode) bool {
	return maxDepthWithBalance(root) != -1
}

func maxDepthWithBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepthWithBalance(root.Left)
	right := maxDepthWithBalance(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
