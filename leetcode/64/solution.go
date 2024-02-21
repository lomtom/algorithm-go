package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:3 ms,击败了93.92% 的Go用户
// 内存消耗:3.3 MB,击败了94.48% 的Go用户
func minPathSum(grid [][]int) int {
	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	m := len(grid)
	n := len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}
	return grid[m-1][n-1]
}

func TestMinPathSum(t *testing.T) {
	collections := []struct {
		input  [][]int
		output int
	}{
		{
			[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			7,
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			12,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, minPathSum(collections[index].input))
	}
}
