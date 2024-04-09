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
			[]int{-2, -2, -1},
			3,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{-2, -1, -1, 1, 2, 3},
			3,
		},
		{
			[]int{-3, -2, -1, 0, 0, 1, 2},
			3,
		},
		{
			[]int{5, 20, 66, 1314},
			4,
		},
		{
			[]int{0, 5, 20, 66, 1314},
			4,
		},
		{
			[]int{-1, 0, 5, 20, 66, 1314},
			4,
		},
		{
			[]int{-1},
			1,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, maximumCount(collections[index].input))
	}
}
