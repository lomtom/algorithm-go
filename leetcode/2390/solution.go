package leetcode

// 执行耗时:20 ms,击败了81.65% 的Go用户
// 内存消耗:6.9 MB,击败了21.52% 的Go用户
func removeStars(s string) string {
	var stack []byte
	for _, v := range s {
		if v == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, byte(v))
		}
	}
	return string(stack)
}
