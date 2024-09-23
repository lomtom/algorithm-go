package leetcode

// 执行耗时:5 ms,击败了79.17% 的Go用户
// 内存消耗:3.6 MB,击败了87.50% 的Go用户
func isToeplitzMatrix(matrix [][]int) bool {
	validate := func(row, col int) bool {
		for row < len(matrix)-1 && col < len(matrix[0])-1 {
			if matrix[row][col] != matrix[row+1][col+1] {
				return false
			}
			row++
			col++
		}
		return true
	}
	for i := 0; i < len(matrix); i++ {
		if !validate(i, 0) {
			return false
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if !validate(0, i) {
			return false
		}
	}
	return true
}
