package leetcode

func getTargetCopy(original, cloned, target *TreeNode) *TreeNode {
	if original == nil {
		return nil
	}
	if original == target {
		return cloned
	}
	if l := getTargetCopy(original.Left, cloned.Left, target); l != nil {
		return l
	}
	return getTargetCopy(original.Right, cloned.Right, target)
}
