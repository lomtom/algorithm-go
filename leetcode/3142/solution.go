package leetcode

// 执行耗时:9 ms,击败了24.49% 的Go用户
// 内存消耗:3.6 MB,击败了100.00% 的Go用户
func satisfiesConditions(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i > 0 && grid[i-1][j] != grid[i][j] {
				return false
			}
			if j > 0 && grid[i][j-1] == grid[i][j] {
				return false
			}
		}
	}
	return true
}
