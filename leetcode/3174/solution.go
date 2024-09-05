package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.2 MB,击败了99.34% 的Go用户
func clearDigits(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
