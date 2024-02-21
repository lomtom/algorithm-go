package leetcode

// 执行用时：8 ms, 在所有 Go 提交中击败了18.84%的用户
// 内存消耗：3.6 MB, 在所有 Go 提交中击败了52.17%的用户
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for i, poppedIndex := 0, 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[poppedIndex] {
			stack = stack[:len(stack)-1]
			poppedIndex++
		}
	}
	return len(stack) == 0
}
