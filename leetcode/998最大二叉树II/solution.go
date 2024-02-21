package leetcode

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val < val {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
