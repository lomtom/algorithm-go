package leetcode

import "strings"

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.2 MB, 在所有 Go 提交中击败了100.00%的用户
func stringMatching(words []string) (ans []string) {
	l := len(words)
	flag := make([]bool, l)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if flag[i] || flag[j] {
				continue
			}
			if strings.Contains(words[i], words[j]) {
				ans = append(ans, words[j])
				flag[j] = true
			}
			if strings.Contains(words[j], words[i]) {
				ans = append(ans, words[i])
				flag[i] = true
			}
		}
	}
	return
}
