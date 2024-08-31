package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.1 MB,击败了100.00% 的Go用户
func canMakeSquare(grid [][]byte) bool {
	for i := range grid {
		for j := range grid[i] {
			if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] == grid[i-1][j] && grid[i][j-1] == grid[i-1][j] {
				return true
			}
			if i-1 >= 0 && j+1 < len(grid[i]) && grid[i-1][j] == grid[i-1][j+1] && grid[i][j+1] == grid[i-1][j+1] {
				return true
			}
			if i+1 < len(grid) && j-1 >= 0 && grid[i][j-1] == grid[i+1][j-1] && grid[i+1][j-1] == grid[i+1][j] {
				return true
			}
			if i+1 < len(grid) && j+1 < len(grid[i]) && grid[i+1][j] == grid[i][j+1] && grid[i+1][j] == grid[i+1][j+1] {
				return true
			}
		}
	}
	return false
}
