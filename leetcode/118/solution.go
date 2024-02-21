package leetcode

// 输入: numRows = 5
// 输出: [
//
//	[1],
//	[1,1],
//	[1,2,1],
//	[1,3,3,1],
//	[1,4,6,4,1]]
//
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.3 MB,击败了22.18% 的Go用户
func generate(numRows int) (result [][]int) {
	for i := 0; i < numRows; i++ {
		var row []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, result[i-1][j]+result[i-1][j-1])
			}
		}
		result = append(result, row)
	}
	return
}
