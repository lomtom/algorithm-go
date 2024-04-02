package leetcode

// 执行耗时:22 ms,击败了38.46% 的Go用户
// 内存消耗:8.7 MB,击败了61.54% 的Go用户
func allPossibleFBT(n int) []*TreeNode {
	var result []*TreeNode
	if n%2 == 0 {
		return result
	}
	if n == 1 {
		result = append(result, &TreeNode{Val: 0})
		return result
	}
	for i := 1; i < n; i += 2 {
		leftSubtrees := allPossibleFBT(i)
		rightSubtrees := allPossibleFBT(n - 1 - i)
		for _, leftSubtree := range leftSubtrees {
			for _, rightSubtree := range rightSubtrees {
				result = append(result, &TreeNode{Val: 0, Left: leftSubtree, Right: rightSubtree})
			}
		}
	}
	return result
}
