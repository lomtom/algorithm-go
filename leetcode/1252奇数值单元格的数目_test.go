package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了90.91%的用户
func oddCells(m int, n int, indices [][]int) (ans int) {
	rows := make([]int, m)
	cols := make([]int, n)
	for index := range indices {
		rows[indices[index][0]] ^= 1
		cols[indices[index][1]] ^= 1
	}
	for _, row := range rows {
		for _, col := range cols {
			ans += (row + col) & 1
		}
	}
	return
}

func TestOddCells(t *testing.T) {
	type input struct {
		m       int
		n       int
		indices [][]int
	}
	collections := []struct {
		input
		output int
	}{
		{
			input{
				2,
				3,
				[][]int{{0, 1}, {1, 1}},
			},
			6,
		},
		{
			input{
				2,
				2,
				[][]int{{1, 1}, {0, 0}},
			},
			0,
		},
	}

	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, oddCells(collections[index].input.m, collections[index].input.n, collections[index].input.indices))
	}
}
