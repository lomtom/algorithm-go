package leetcode

import "unicode"

// 执行耗时:5 ms,击败了80.81% 的Go用户
// 内存消耗:6.2 MB,击败了40.40% 的Go用户
func letterCasePermutation(s string) []string {
	result := make([]string, 0)
	dfs([]rune(s), 0, len(s), &result)
	return result
}

func dfs(s []rune, index int, l int, result *[]string) {
	if index == l {
		*result = append(*result, string(s))
		return
	}
	dfs(s, index+1, l, result)
	if unicode.IsLetter(s[index]) {
		s[index] = s[index] ^ 32
		dfs(s, index+1, l, result)
	}
}
