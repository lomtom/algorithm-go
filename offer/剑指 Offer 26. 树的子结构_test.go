package offer

import (
	"algorithm/leetcode"
	"testing"
)

func isSubStructure(A *leetcode.TreeNode, B *leetcode.TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(A *leetcode.TreeNode, B *leetcode.TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
}

// 测试用例
func Test_isSubStructure(t *testing.T) {
	// 示例 1
	A := &leetcode.TreeNode{
		Val: 3,
		Left: &leetcode.TreeNode{
			Val: 4,
			Left: &leetcode.TreeNode{
				Val: 1,
			},
			Right: &leetcode.TreeNode{
				Val: 2,
			},
		},
		Right: &leetcode.TreeNode{
			Val: 5,
		},
	}
	B := &leetcode.TreeNode{
		Val: 4,
		Left: &leetcode.TreeNode{
			Val: 1,
		},
	}
	t.Log(isSubStructure(A, B))
}
