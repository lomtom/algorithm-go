package leetcode

import (
	"fmt"
	"testing"
)

func coinChange(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 0; i <= amount; i++ {
		for index := range coins {
			if i-coins[index] < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[index]]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 0))
	fmt.Println(coinChange([]int{1, 2, 5}, 1))
	fmt.Println(coinChange([]int{1, 2, 5}, 2))
	fmt.Println(coinChange([]int{1, 2, 5}, 3))
	fmt.Println(coinChange([]int{1, 2, 5}, 4))
	fmt.Println(coinChange([]int{1, 2, 5}, 5))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

// 1 2 5
// 0 0
// 1 1
// 2 1
// 3 2
// 4 2
// 5 1
// 6 2
// 7 2
// 8 3
// 9 3
// 10 2
