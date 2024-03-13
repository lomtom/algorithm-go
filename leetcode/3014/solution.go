package leetcode

// 执行耗时:2 ms,击败了22.43% 的Go用户
// 内存消耗:1.9 MB,击败了81.31% 的Go用户
func minimumPushes(word string) int {
	l := len(word)
	if l <= 8 {
		return l - 0
	}
	if l <= 16 {
		return 2*l - 8
	}
	if l <= 24 {
		return 3*l - 24
	}
	return 4*l - 48
}
