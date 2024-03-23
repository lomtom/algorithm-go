package leetcode

// 执行耗时:1 ms,击败了30.43% 的Go用户
// 内存消耗:2 MB,击败了17.39% 的Go用户
func distinctIntegers(n int) int {
	if n > 1 {
		return n - 1
	}
	return n
}
