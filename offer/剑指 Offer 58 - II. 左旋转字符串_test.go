package offer

import "testing"

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func TestReverseLeftWords(t *testing.T) {
	t.Log(reverseLeftWords("abcdefg", 2))
	t.Log(reverseLeftWords("lrloseumgh", 6))
}
