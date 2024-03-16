package leetcode

// 执行耗时:138 ms,击败了63.89% 的Go用户
// 内存消耗:8.5 MB,击败了19.44% 的Go用户
func maxMoves(grid [][]int) (result int) {
	rows, cols := len(grid), len(grid[0])
	q := make([][2]int, 0)
	for i := 0; i < rows; i++ {
		q = append(q, [2]int{i, 0})
	}
	for len(q) != 0 {
		num := q[0]
		q = q[1:]
		if grid[num[0]][num[1]] == 0 {
			continue
		}
		q1 := make([][2]int, 0)
		if num[0]-1 >= 0 && num[1]+1 < cols && grid[num[0]-1][num[1]+1] > grid[num[0]][num[1]] {
			q1 = append(q1, [2]int{num[0] - 1, num[1] + 1})
		}
		if num[1]+1 < cols && grid[num[0]][num[1]+1] > grid[num[0]][num[1]] {
			q1 = append(q1, [2]int{num[0], num[1] + 1})
		}
		if num[0]+1 < rows && num[1]+1 < cols && grid[num[0]+1][num[1]+1] > grid[num[0]][num[1]] {
			q1 = append(q1, [2]int{num[0] + 1, num[1] + 1})
		}
		if num[1] > result {
			result = num[1]
		}
		grid[num[0]][num[1]] = 0
		q = append(q, q1...)
	}
	return
}
