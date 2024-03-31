package leetcode

var direction = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 执行耗时:6 ms,击败了77.27% 的Go用户
// 内存消耗:6.3 MB,击败了13.64% 的Go用户
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) (result [][]int) {
	var count = rows * cols
	var l = 1
	for count > 0 {
		for index := range direction {
			for i := 0; i < l; i++ {
				if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
					result = append(result, []int{rStart, cStart})
					count--
				}
				rStart += direction[index][0]
				cStart += direction[index][1]
			}
			if index%2 == 1 {
				l++
			}
		}
	}
	return
}
