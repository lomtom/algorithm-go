package leetcode

import "math"

// 9	5	7	3
// 8	9	6	1
// 6	7	14	3
// 2	5	3	1

// 找到当前值右下角矩阵的最大值

// 执行耗时:115 ms,击败了60.51% 的Go用户
// 内存消耗:8.4 MB,击败了30.77% 的Go用户
func maxScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	ans := grid[0][1] - grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var temp = grid[i][j]
			if i > 0 {
				temp = min(temp, dp[i-1][j])
				ans = max(ans, grid[i][j]-dp[i-1][j])
			}
			if j > 0 {
				temp = min(temp, dp[i][j-1])
				ans = max(ans, grid[i][j]-dp[i][j-1])
			}
			dp[i][j] = temp
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
