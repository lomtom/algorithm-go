package leetcode

var m = map[uint8]uint8{
	')': '(',
	'}': '{',
	']': '[',
}

// 执行耗时:1 ms,击败了12.88% 的Go用户
// 内存消耗:2 MB,击败了15.79% 的Go用户
func isValid(s string) bool {
	stack := make([]uint8, 0)
	for index := range s {
		if _, ok := m[s[index]]; !ok {
			stack = append(stack, s[index])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[index]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
