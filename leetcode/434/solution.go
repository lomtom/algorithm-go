package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:3.5 MB,击败了34.00% 的Go用户
func countSegments(s string) int {
	var count int = 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && (i == 0 || s[i-1] == ' ') {
			count++
		}
	}
	return count
}
