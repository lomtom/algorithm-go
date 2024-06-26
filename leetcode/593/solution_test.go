package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func Test(t *testing.T) {
	type input struct {
		p1 []int
		p2 []int
		p3 []int
		p4 []int
	}
	collections := []struct {
		input
		output bool
	}{
		{
			input{
				[]int{1, 1},
				[]int{0, 1},
				[]int{1, 2},
				[]int{0, 0},
			},
			false,
		},
		{
			input{
				[]int{0, 0},
				[]int{1, 1},
				[]int{1, 0},
				[]int{0, 1},
			},
			true,
		},
		{
			input{
				[]int{0, 0},
				[]int{1, 1},
				[]int{1, 0},
				[]int{0, 12},
			},
			false,
		},
		{
			input{
				[]int{1, 0},
				[]int{-1, 0},
				[]int{0, 1},
				[]int{0, -1},
			},
			true,
		},
		{
			input{
				[]int{0, 0},
				[]int{0, 0},
				[]int{0, 0},
				[]int{0, 0},
			},
			false,
		},
		{
			input{
				[]int{0, 1},
				[]int{1, 2},
				[]int{0, 2},
				[]int{0, 0},
			},
			false,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, validSquare(collections[index].input.p1, collections[index].input.p2, collections[index].input.p3, collections[index].input.p4))
	}
}
