package leetcode

// 执行耗时:16 ms,击败了16.67% 的Go用户
// 内存消耗:6.3 MB,击败了50.00% 的Go用户
func maximumSubsequenceCount(text string, pattern string) int64 {
	var count [2]int
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[0] {
			count[0]++
		}
		if text[i] == pattern[1] {
			count[1]++
		}
	}
	var ans int64
	if count[0] > count[1] {
		ans += int64(count[0])
	} else {
		ans += int64(count[1])
	}
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[1] {
			count[1]--
		}
		if text[i] == pattern[0] {
			count[0]--
			ans += int64(count[1])
		}
	}
	return ans
}
