package leetcode

// 执行耗时:43 ms,击败了13.04% 的Go用户
// 内存消耗:6.7 MB,击败了13.04% 的Go用户
func countConsistentStrings(allowed string, words []string) int {
	m := make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		m[allowed[i]] = true
	}
	res := 0
	for i := 0; i < len(words); i++ {
		var flag = true
		for j := 0; j < len(words[i]); j++ {
			flag = flag && m[words[i][j]]
		}
		if flag {
			res++
		}
	}
	return res
}
