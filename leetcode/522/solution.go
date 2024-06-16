package leetcode

// judge whether s is a subsequence of t
func isSubseq(s, t string) bool {
	i := 0
	for _, c := range t {
		if s[i] == byte(c) {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}

// judge whether strs[index] is a subsequence of other strings
func hasSubSeq(index int, strs []string) bool {
	for j, t := range strs {
		if j != index && isSubseq(strs[index], t) {
			return false
		}
	}
	return true
}

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.1 MB,击败了88.89% 的Go用户
func findLUSlength(strs []string) int {
	ans := -1
	for i, s := range strs {
		if len(s) <= ans {
			continue
		}
		if hasSubSeq(i, strs) {
			ans = len(strs[i])
		}
	}
	return ans
}
