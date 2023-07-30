package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

func shiftGrid(grid [][]int, k int) [][]int {
	rows, cols := len(grid), len(grid[0])
	ans := make([][]int, rows)
	for i := range ans {
		ans[i] = make([]int, cols)
	}
	for i, row := range grid {
		for j, v := range row {
			index := (i*cols + j + k) % (rows * cols)
			ans[index/cols][index%cols] = v
		}
	}
	return ans
}

func TestShiftGrid(t *testing.T) {
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
