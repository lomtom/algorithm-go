package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func Test(t *testing.T) {
	collections := []struct {
		input  string
		output string
	}{
		{
			"  hello world  ",
			"world hello",
		},
		{
			"the sky is blue",
			"blue is sky the",
		},
		{
			"a good   example",
			"example good a",
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, reverseWords(collections[index].input))
	}
}
