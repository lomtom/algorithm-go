package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func Test(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{200, 3, 140, 20, 10},
			340,
		},
		{
			[]int{2, 3, 2},
			3,
		},
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{[]int{1, 2, 3},
			3,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, robII(collections[index].input))
	}
}
