package leetcode

var direction = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 执行耗时:1 ms,击败了25.53% 的Go用户
// 内存消耗:2 MB,击败了21.59% 的Go用户
func spiralOrder(matrix [][]int) (result []int) {
	rows := len(matrix)
	cols := len(matrix[0])
	var row, col = 0, 0
	var directionCount = 1
	var count = 0
	var index = 0
	for count < rows*cols {
		result = append(result, matrix[row][col])
		count++
		if row+direction[index][0]+direction[index][0]*directionCount/4 < 0 || row+direction[index][0]+direction[index][0]*directionCount/4 >= rows ||
			col+direction[index][1]+direction[index][1]*directionCount/4 < 0 || col+direction[index][1]+direction[index][1]*directionCount/4 >= cols {
			index = (index + 1) % 4
			directionCount++
		}
		row += direction[index][0]
		col += direction[index][1]
	}
	return
}
