package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

// dfs 超时
//func lenLongestFibSubseq(arr []int) int {
//	var result int
//	var dfs func(arr []int, start, l, num1, num2, num int)
//	dfs = func(arr []int, start, l, num1, num2, num int) {
//		if result < num && num > 2  {
//			result = num
//		}
//		if start >= l {
//			return
//		}
//		for index := start; index < l; index++ {
//			if num1+num2 == arr[index] || num1 == 0 || num2 == 0 {
//				num++
//				dfs(arr, index+1, l, num2, arr[index], num)
//				num--
//			}
//		}
//	}
//	l := len(arr)
//	for index := 0; index < l; index++ {
//		dfs(arr, 0, l, 0, 0, 0)
//	}
//	return result
//}

func lenLongestFibSubseq(arr []int) (ans int) {
	l := len(arr)
	m := make(map[int]int)
	dp := make([][]int, l)
	for index := 0; index < l; index++ {
		m[arr[index]] = index
		dp[index] = make([]int, l)
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	for raw := l - 1; raw >= 0; raw-- {
		for col := l - 1; col > raw; col-- {
			if index, ok := m[arr[col]+arr[raw]]; ok {
				dp[raw][col] = dp[col][index] + 1
				ans = max(ans, dp[raw][col])
			}
		}
	}
	if ans == 0 {
		return
	}
	return ans + 2
}
func TestLenLongestFibSubseq(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			5,
		},
		{
			[]int{1, 3, 7, 11, 12, 14, 18},
			3,
		},
		{
			// 3,15,18,32,50
			[]int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50},
			5,
		},
		{
			[]int{1, 3, 5},
			0,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, lenLongestFibSubseq(collections[index].input))
	}
}
