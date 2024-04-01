package leetcode

import "strings"

// 执行耗时:4 ms,击败了12.50% 的Go用户
// 内存消耗:3 MB,击败了5.36% 的Go用户
func isValidSerialization(preorder string) bool {
	ss := strings.Split(preorder, ",")
	stack := make([]string, 0)
	for index := range ss {
		stack = append(stack, ss[index])
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" {
			stack = stack[:len(stack)-2]
			if stack[len(stack)-1] == "#" {
				return false
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, "#")
		}
	}
	return len(stack) == 1 && stack[0] == "#"
}
