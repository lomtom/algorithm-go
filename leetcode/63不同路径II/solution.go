package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.3 MB,击败了75.28% 的Go用户
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] += dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func TestUniquePathsWithObstacles(t *testing.T) {
	collections := []struct {
		input  [][]int
		output int
	}{
		{
			[][]int{
				{0, 1, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			0,
		},
		{
			[][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			},
			0,
		},
		{
			[][]int{
				{0, 0},
				{0, 1},
			},
			0,
		},
		{
			[][]int{
				{1, 0},
				{0, 0},
			},
			0,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			2,
		},
		{
			[][]int{
				{0, 1}, {0, 0},
			},
			1,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, uniquePathsWithObstacles(collections[index].input))
	}
}
