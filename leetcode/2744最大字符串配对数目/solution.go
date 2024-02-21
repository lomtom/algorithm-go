package leetcode

import (
	"fmt"
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:4 ms,击败了48.00% 的Go用户
// 内存消耗:3.3 MB,击败了12.00% 的Go用户
func maximumNumberOfStringPairs(words []string) int {
	m := make(map[string]bool)
	total := 0
	for i := range words {
		if _, ok := m[fmt.Sprintf("%s%s", string(words[i][1]), string(words[i][0]))]; ok {
			total++
			m[words[i]] = false
		} else {
			m[words[i]] = true
		}
	}
	return total
}

func TestMaximumNumberOfStringPairs(t *testing.T) {
	collections := []struct {
		input  []string
		output int
	}{
		{
			[]string{"cd", "ac", "dc", "ca", "zz"},
			2,
		},
		{[]string{"ab", "ba", "cc"},
			1,
		},
		{[]string{"aa", "ab"},
			0,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, maximumNumberOfStringPairs(collections[index].input))
	}
}
