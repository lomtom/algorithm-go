package leetcode

func scoreOfParentheses(s string) int {
	n := len(s)
	if n == 2 {
		return 1
	}
	for i, bal := 0, 0; ; i++ {
		if s[i] == '(' {
			bal++
		} else {
			bal--
			if bal == 0 {
				if i == n-1 {
					return 2 * scoreOfParentheses(s[1:n-1])
				}
				return scoreOfParentheses(s[:i+1]) + scoreOfParentheses(s[i+1:])
			}
		}
	}
}

// 栈
//func scoreOfParentheses(s string) int {
//	st := []int{0}
//	for _, c := range s {
//		if c == '(' {
//			st = append(st, 0)
//		} else {
//			v := st[len(st)-1]
//			st = st[:len(st)-1]
//			st[len(st)-1] += max(2*v, 1)
//		}
//	}
//	return st[0]
//}
//
//func max(a, b int) int {
//	if b > a {
//		return b
//	}
//	return a
//}
