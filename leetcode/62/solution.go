package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了61.56% 的Go用户
func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	// init dp
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func TestUniquePaths(t *testing.T) {
	collections := []struct {
		input  [2]int
		output int
	}{
		{
			[2]int{3, 7},
			28,
		},
		{[2]int{3, 3},
			6,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, uniquePaths(collections[index].input[0], collections[index].input[1]))
	}
}
