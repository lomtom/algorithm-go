package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func Test(t *testing.T) {
	type input struct {
		grid [][]int
		k    int
	}
	collections := []struct {
		input
		output [][]int
	}{
		{
			input: input{
				[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				1,
			},
			output: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
	}

	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, shiftGrid(collections[index].input.grid, collections[index].input.k))
	}
}
