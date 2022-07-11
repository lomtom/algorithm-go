package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"math"
	"testing"
)

// bfs 超时
//func cherryPickup(grid [][]int) (ans int) {
//	// 右、下、左、上
//	goStep := [][2]int{{0, 1}, {1, 0}}
//	backStep := [][2]int{{0, -1}, {-1, 0}}
//	rawLen := len(grid)
//	colLen := len(grid[0])
//	if rawLen == 1 {
//		return grid[0][0]
//	}
//	var goDfs func(rawNow, colNow, result int)
//	var backDfs func(rawNow, colNow, result int)
//	goDfs = func(rawNow, colNow, result int) {
//		if rawNow == rawLen-1 && colNow == colLen-1 {
//			backDfs(rawNow, colNow, result)
//		}
//		for raw := rawNow; raw < rawLen; raw++ {
//			for col := colNow; col < colLen; col++ {
//				flag := false
//				for index := range goStep {
//					if raw+goStep[index][0] < rawLen && col+goStep[index][1] < colLen {
//						num := grid[raw+goStep[index][0]][col+goStep[index][1]]
//						if num == 1 || num == 0 {
//							temp := num
//							grid[raw+goStep[index][0]][col+goStep[index][1]] = 0
//							goDfs(raw+goStep[index][0], col+goStep[index][1], result+temp)
//							grid[raw+goStep[index][0]][col+goStep[index][1]] = temp
//							flag = true
//						} else {
//							continue
//						}
//					} else {
//						flag = true
//					}
//				}
//				// 无路可走
//				if !flag {
//					ans = 0
//					return
//				}
//			}
//		}
//	}
//	backDfs = func(rawNow, colNow, result int) {
//		if rawNow == 0 && colNow == 0 && result > ans {
//			ans = result
//			return
//		}
//		for raw := rawNow; raw >= 0; raw-- {
//			for col := colNow; col >= 0; col-- {
//				for index := range backStep {
//					if raw+backStep[index][0] >= 0 && col+backStep[index][1] >= 0 {
//						num := grid[raw+backStep[index][0]][col+backStep[index][1]]
//						if num == 1 || num == 0 {
//							temp := num
//							grid[raw+backStep[index][0]][col+backStep[index][1]] = 0
//							backDfs(raw+backStep[index][0], col+backStep[index][1], result+temp)
//							grid[raw+backStep[index][0]][col+backStep[index][1]] = temp
//						} else {
//							continue
//						}
//					}
//				}
//			}
//		}
//	}
//	goDfs(0, 0, 0)
//	return
//}

//执行用时：4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：4.1 MB, 在所有 Go 提交中击败了100.00%的用户
func cherryPickup(grid [][]int) int {
	rawLen := len(grid)
	colLen := len(grid[0])
	dp := make([][]int, rawLen)
	for raw := 0; raw < rawLen; raw++ {
		for col := 0; col < colLen; col++ {
			dp[raw][col] = math.MinInt32
		}
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	dp[0][0] = grid[0][0]
	for k := 1; k < rawLen*2-1; k++ {
		for x1 := min(k, rawLen-1); x1 >= max(k-rawLen+1, 0); x1-- {
			for x2 := min(k, rawLen-1); x2 >= x1; x2-- {
				y1, y2 := k-x1, k-x2
				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					dp[x1][x2] = math.MinInt32
					continue
				}
				// 都往右
				res := dp[x1][x2]
				// 往下，往右
				if x1 > 0 {
					res = max(res, dp[x1-1][x2])
				}
				// 往右，往下
				if x2 > 0 {
					res = max(res, dp[x1][x2-1])
				}
				// 都往下
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[x1-1][x2-1])
				}
				res += grid[x1][y1]
				// 避免重复摘同一个樱桃
				if x2 != x1 {
					res += grid[x2][y2]
				}
				dp[x1][x2] = res
			}
		}
	}
	return max(dp[rawLen-1][rawLen-1], 0)
}

func TestCherryPickup(t *testing.T) {
	collections := []struct {
		input  [][]int
		output int
	}{
		{
			[][]int{
				{0, 1, -1},
				{1, 0, -1},
				{1, 1, 1},
			},
			5,
		},
		{
			[][]int{
				{0, -1},
				{-1, 0},
			},
			0,
		},
		{
			[][]int{{1}},
			1,
		},
		{
			[][]int{{0}},
			0,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, cherryPickup(collections[index].input))
	}
}
