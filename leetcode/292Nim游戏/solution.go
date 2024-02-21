package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.8 MB,击败了74.14% 的Go用户
func canWinNim(n int) bool {
	return n%4 != 0
}
