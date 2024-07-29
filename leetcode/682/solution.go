package leetcode

import "strconv"

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.6 MB,击败了66.67% 的Go用户
func calPoints(operations []string) int {
	var stack []int
	for _, op := range operations {
		switch op {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
	}
	var ret int
	for _, v := range stack {
		ret += v
	}
	return ret
}
