package leetcode

// 执行耗时:126 ms,击败了54.10% 的Go用户
// 内存消耗:3.9 MB,击败了55.40% 的Go用户
func exist(board [][]byte, word string) bool {
	var visited [][]bool
	var m, n, l = len(board), len(board[0]), len(word)
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}
		if visited[i][j] || board[i][j] != word[index] {
			return false
		}
		if index == l-1 {
			return true
		}
		visited[i][j] = true
		flag1 := dfs(i, j+1, index+1)
		flag2 := dfs(i+1, j, index+1)
		flag3 := dfs(i, j-1, index+1)
		flag4 := dfs(i-1, j, index+1)
		visited[i][j] = false
		return flag1 || flag2 || flag3 || flag4
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
