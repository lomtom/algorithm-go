package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func countGoodSubstrings(s string) (ans int) {
	for i := 0; i <= len(s)-3; i++ {
		if s[i] != s[i+1] && s[i+1] != s[i+2] && s[i] != s[i+2] {
			ans++
		}
	}
	return ans
}

func TestCountGoodSubstrings(t *testing.T) {
	collections := []struct {
		input  string
		output int
	}{
		{
			"xyzzaz",
			1,
		},
		{
			"aababcabc",
			4,
		},
	}

	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, countGoodSubstrings(collections[index].input))
	}
}
