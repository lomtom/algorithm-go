package leetcode

// 执行耗时:1 ms,击败了40.84% 的Go用户
// 内存消耗:2.4 MB,击败了28.50% 的Go用户
func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
