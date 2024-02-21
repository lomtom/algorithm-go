package leetcode

import "strings"

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了27.78%的用户
func reformatNumber(number string) string {
	number = strings.Replace(number, " ", "", -1)
	number = strings.Replace(number, "-", "", -1)
	var ans []string
	for len(number) > 4 {
		ans = append(ans, number[:3])
		number = number[3:]
	}
	if len(number) == 4 {
		ans = append(ans, number[:2])
		number = number[2:]
	}
	ans = append(ans, number)
	return strings.Join(ans, "-")
}
