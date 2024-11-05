package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:3.9 MB,击败了6.90% 的Go用户
func losingPlayer(x int, y int) string {
	opt := min(x, y/4)
	if opt%2 == 0 {
		return "Bob"
	} else {
		return "Alice"
	}
}
