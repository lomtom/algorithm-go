package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	inStack := [26]bool{}
	var stack []uint8
	for index := 0; index < len(s); index++ {
		if !inStack[s[index]-'a'] {
			for len(stack) > 0 && stack[len(stack)-1] > s[index] && left[stack[len(stack)-1]-'a'] != 0 {
				inStack[stack[len(stack)-1]-'a'] = false
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, s[index])
			inStack[s[index]-'a'] = true
		}
		left[s[index]-'a']--
	}
	return string(stack)
}

//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:1.9 MB,击败了82.21% 的Go用户
func TestRemoveDuplicateLetters(t *testing.T) {
	collections := []struct {
		input  string
		output string
	}{
		{
			"bcabc",
			"abc",
		},
		{"cbacdcbc",
			"acdb",
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, removeDuplicateLetters(collections[index].input))
	}
}
