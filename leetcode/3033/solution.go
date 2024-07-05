package leetcode

// 执行耗时:9 ms,击败了30.00% 的Go用户
// 内存消耗:4 MB,击败了43.33% 的Go用户
func modifiedMatrix(matrix [][]int) [][]int {
	var m, n = len(matrix), len(matrix[0])
	var max = make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[j][i] >= max[i] {
				max[i] = matrix[j][i]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[j][i] == -1 {
				matrix[j][i] = max[i]
			}
		}
	}
	return matrix
}
