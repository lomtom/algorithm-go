package leetcode

// 执行耗时:11 ms,击败了45.00% 的Go用户
// 内存消耗:6.3 MB,击败了12.50% 的Go用户
func longestContinuousSubstring(s string) (ans int) {
	var count int = 1
	ans = count
	for i := 1; i < len(s); i++ {
		if int(s[i])-int(s[i-1]) == 1 {
			count++
		} else {
			count = 1
		}
		ans = max(ans, count)
	}
	return ans
}
