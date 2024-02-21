package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
	"unicode"
)

func reformat(s string) string {
	digitNum := 0
	for _, c := range s {
		if unicode.IsDigit(c) {
			digitNum++
		}
	}
	alphaNum := len(s) - digitNum
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	if abs(digitNum-alphaNum) > 1 {
		return ""
	}
	flag := digitNum < alphaNum
	var ans []rune
	for i, j := 0, 0; i < len(s) || j < len(s); {
		flag = !flag
		if flag {
			for i < len(s) && !unicode.IsDigit(rune(s[i])) {
				i++
			}
			if i == len(s) {
				continue
			}
			ans = append(ans, rune(s[i]))
			i++
		} else {
			for j < len(s) && unicode.IsDigit(rune(s[j])) {
				j++
			}
			if j == len(s) {
				continue
			}
			ans = append(ans, rune(s[j]))
			j++
		}
	}
	return string(ans)
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：3.3 MB, 在所有 Go 提交中击败了31.43%的用户
func TestReformat(t *testing.T) {
	collections := []struct {
		input  string
		output string
	}{
		{
			"a0b1c2",
			"a0b1c2",
		},
		{
			"leetcode",
			"",
		},
		{
			"1229857369",
			"",
		},
		{
			"covid2019",
			"c2o0v1i9d",
		},
		{
			"ab123",
			"1a2b3",
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, reformat(collections[index].input))
	}
}
