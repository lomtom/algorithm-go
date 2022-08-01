package leetcode

import (
	"strings"
)

//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2 MB,击败了94.12% 的Go用户
func generateTheString(n int) (ans string) {
	if n%2 == 1 {
		return strings.Repeat("a", n)
	}
	return strings.Repeat("a", n-1) + "b"
}
