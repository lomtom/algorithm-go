package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"strings"
	"testing"
)

func removeKdigits(num string, k int) string {
	stack := make([]uint8, 0)
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if len(ans) == 0 {
		return "0"
	}
	return ans
}

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:4.5 MB,击败了78.51% 的Go用户
func TestRemoveKdigits(t *testing.T) {
	type input struct {
		num string
		k   int
	}
	collections := []struct {
		input
		output string
	}{
		{
			input{
				"9",
				1,
			},
			"0",
		},
		{
			input{
				"1432219",
				3,
			},
			"1219",
		},
		{
			input{
				"10200",
				1,
			},
			"200",
		},
		{
			input{
				"10",
				2,
			},
			"0",
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, removeKdigits(collections[index].input.num, collections[index].input.k))
	}
}
