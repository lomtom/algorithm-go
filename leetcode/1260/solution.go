package leetcode

func shiftGrid(grid [][]int, k int) [][]int {
	rows, cols := len(grid), len(grid[0])
	ans := make([][]int, rows)
	for i := range ans {
		ans[i] = make([]int, cols)
	}
	for i, row := range grid {
		for j, v := range row {
			index := (i*cols + j + k) % (rows * cols)
			ans[index/cols][index%cols] = v
		}
	}
	return ans
}
