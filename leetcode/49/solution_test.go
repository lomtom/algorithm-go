package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(groupAnagrams([]string{""}))
	t.Log(groupAnagrams([]string{"a"}))
}
