package leetcode

// 执行耗时:183 ms,击败了89.70% 的Go用户
// 内存消耗:24.4 MB,击败了72.73% 的Go用户
func countSubmatrices(grid [][]int, k int) int {
	var dp = make([][]int, len(grid))
	l := len(grid[0])
	dp[0] = make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = grid[0][i] + dp[0][i-1]
		}
	}
	for i := 1; i < len(grid); i++ {
		dp[i] = make([]int, l)
		for j := range grid[i] {
			if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
			}
		}
	}
	var count int
	for i := range dp {
		for j := range dp[i] {
			if dp[i][j] <= k {
				count++
			}
		}
	}
	return count
}
