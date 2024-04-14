package leetcode

import "strings"

// 执行耗时:1 ms,击败了29.32% 的Go用户
// 内存消耗:2.3 MB,击败了5.12% 的Go用户
func lengthOfLastWord(s string) int {
	ss := strings.Split(strings.TrimRight(s, " "), " ")
	return len(ss[len(ss)-1])
}
