package leetcode

import "math"

// 执行耗时:3 ms,击败了48.86% 的Go用户
// 内存消耗:4.1 MB,击败了15.91% 的Go用户
func maxAncestorDiff1(root *TreeNode) (result int) {
	var dfs func(root *TreeNode) (min, max int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return math.MaxInt, math.MinInt
		}
		max, min := root.Val, root.Val
		lMin, lMax := dfs(root.Left)
		rMin, rMax := dfs(root.Right)
		max = maxFunc(maxFunc(max, lMax), rMax)
		min = minFunc(minFunc(min, lMin), rMin)
		result = maxFunc(result, maxFunc(root.Val-min, max-root.Val))
		return min, max
	}
	dfs(root)
	return result
}

// 执行耗时:3 ms,击败了48.86% 的Go用户
// 内存消耗:4.1 MB,击败了9.09% 的Go用户
func maxAncestorDiff(root *TreeNode) int {
	var dfs func(root *TreeNode, min, max int) int
	dfs = func(root *TreeNode, min, max int) int {
		if root == nil {
			return max - min
		}
		min = minFunc(min, root.Val)
		max = maxFunc(max, root.Val)
		return maxFunc(dfs(root.Left, min, max), dfs(root.Right, min, max))
	}
	return dfs(root, root.Val, root.Val)
}

func minFunc(value1, value2 int) int {
	if value2 < value1 {
		return value2
	}
	return value1
}

func maxFunc(value1, value2 int) int {
	if value2 > value1 {
		return value2
	}
	return value1
}
