package leetcode

// 执行耗时:264 ms,击败了54.82% 的Go用户
// 内存消耗:22.7 MB,击败了63.70% 的Go用户
func numberOfRightTriangles(grid [][]int) int64 {
	var rowCount = make([]int, len(grid))
	var colCount = make([]int, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			rowCount[i] += grid[i][j]
			colCount[j] += grid[i][j]
		}
	}
	var res int64
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				res += (int64(rowCount[i]) - 1) * (int64(colCount[j]) - 1)
			}
		}
	}
	return res
}
