package leetcode

import "strconv"

// 执行耗时:8 ms,击败了85.92% 的Go用户
// 内存消耗:4.3 MB,击败了32.72% 的Go用户
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i <= (l-1)/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

// 执行耗时:14 ms,击败了45.76% 的Go用户
// 内存消耗:4.12 MB,击败了88.07% 的Go用户
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	var x1, y = x, 0
	for x1 != 0 {
		y = y*10 + x1%10
		x1 /= 10
	}
	return x == y
}
