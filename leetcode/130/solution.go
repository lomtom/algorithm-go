package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:7.9 MB,击败了94.79% 的Go用户
func solve(board [][]byte) {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'U'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := range board {
		for j := range board[i] {
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1) && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'U' {
				board[i][j] = 'O'
			}
		}
	}
	return
}
