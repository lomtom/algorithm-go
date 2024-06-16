package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(findLUSlength("aba", "cdc"))
	t.Log(findLUSlength("aaa", "bbb"))
	t.Log(findLUSlength("aaa", "aaa"))
	t.Log(findLUSlength("abc", "aebdc"))
}
