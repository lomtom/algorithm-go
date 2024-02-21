package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 7,1,5,3,6,4
// 执行耗时:94 ms,击败了85.62% 的Go用户
// 内存消耗:7.7 MB,击败了80.61% 的Go用户
func maxProfit(prices []int) int {
	var min = prices[0]
	var max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	var dp = make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		dp[i] = max(dp[i-1], prices[i]-min)
	}
	return dp[len(prices)-1]
}

func TestMaxProfit(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{[]int{7, 6, 4, 3, 1},
			0,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, maxProfit(collections[index].input))
	}
}
