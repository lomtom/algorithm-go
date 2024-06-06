package leetcode

// 执行耗时:7 ms,击败了100.00% 的Go用户
// 内存消耗:6.2 MB,击败了70.83% 的Go用户
func minimumSteps(s string) (result int64) {
	left, right := 0, len(s)-1
	for left < right {
		for left < len(s) && s[left] == '0' {
			left++
		}
		for right >= 0 && s[right] == '1' {
			right--
		}
		if left < right {
			result += int64(right - left)
		}
		left++
		right--
	}
	return
}
