package leetcode

// 执行耗时:16 ms,击败了62.30% 的Go用户
// 内存消耗:6.2 MB,击败了98.97% 的Go用户
func searchMatrix(matrix [][]int, target int) bool {
	if matrix[0][0] > target || matrix[len(matrix)-1][len(matrix[0])-1] < target {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		for row < rows && col >= 0 && matrix[row][col] < target {
			row++
		}
		for row < rows && col >= 0 && matrix[row][col] > target {
			col--
		}
	}
	return false
}
