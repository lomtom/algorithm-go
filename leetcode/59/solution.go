package leetcode

var direction = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 执行耗时:1 ms,击败了29.01% 的Go用户
// 内存消耗:2.2 MB,击败了15.94% 的Go用户
func generateMatrix(n int) [][]int {
	var count = 0
	var result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	var directionCount = 1
	var index = 0
	var row, col = 0, 0
	for count < n*n {
		result[row][col] = count + 1
		count++
		if row+direction[index][0]+direction[index][0]*directionCount/4 < 0 || row+direction[index][0]+direction[index][0]*directionCount/4 >= n ||
			col+direction[index][1]+direction[index][1]*directionCount/4 < 0 || col+direction[index][1]+direction[index][1]*directionCount/4 >= n {
			index = (index + 1) % 4
			directionCount++
		}
		row += direction[index][0]
		col += direction[index][1]
	}
	return result
}
