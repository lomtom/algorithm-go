package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:4.5 MB,击败了91.88% 的Go用户
func generateParenthesis(n int) []string {
	var res []string
	var dfs func(bytes []byte, leftNumber, rightNumber int)
	dfs = func(bytes []byte, leftNumber, rightNumber int) {
		if leftNumber == n && rightNumber == n {
			res = append(res, string(bytes))
			return
		}
		if leftNumber < n {
			dfs(append(bytes, '('), leftNumber+1, rightNumber)
		}
		if rightNumber < leftNumber {
			dfs(append(bytes, ')'), leftNumber, rightNumber+1)
		}
	}
	dfs([]byte{}, 0, 0)
	return res
}
