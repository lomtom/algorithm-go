package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(findLUSlength([]string{"aabbcc", "aabbcc", "cb", "abc"}))
	t.Log(findLUSlength([]string{"aba", "cdc", "eae"}))
	t.Log(findLUSlength([]string{"aaa", "aaa", "aaa"}))
}
