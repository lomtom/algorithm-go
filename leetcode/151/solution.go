package leetcode

import "strings"

// 执行耗时:2 ms,击败了49.98% 的Go用户
// 内存消耗:3 MB,击败了60.10% 的Go用户
func reverseWords(s string) string {
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	var left = len(s) - 1
	var result []byte
	for left >= 0 {
		for left >= 0 && s[left] == ' ' {
			left--
		}
		right := left
		for left >= 0 && s[left] != ' ' {
			left--
		}
		result = append(result, s[left+1:right+1]...)
		if left >= 0 {
			result = append(result, ' ')
		}
	}
	return string(result)
}
