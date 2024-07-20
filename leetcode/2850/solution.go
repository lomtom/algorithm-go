package leetcode

import "math"

// 执行耗时:108 ms,击败了16.67% 的Go用户
// 内存消耗:2.7 MB,击败了50.00% 的Go用户
func minimumMoves(grid [][]int) int {
	var num1 [][]int
	var num2 [][]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 1 {
				num1 = append(num1, []int{i, j, grid[i][j]})
			} else if grid[i][j] == 0 {
				num2 = append(num2, []int{i, j, grid[i][j]})
			}
		}
	}
	var minCount = math.MaxInt
	var dfs func(count, zeroCount int)
	dfs = func(count, zeroCount int) {
		if zeroCount == 0 {
			minCount = min(minCount, count)
			return
		}
		for index1 := 0; index1 < len(num1); index1++ {
			for index2 := 0; index2 < len(num2); index2++ {
				if num1[index1][2] > 1 && num2[index2][2] == 0 {
					num1[index1][2]--
					num2[index2][2]++
					dfs(count+abs(num1[index1][0]-num2[index2][0])+abs(num1[index1][1]-num2[index2][1]), zeroCount-1)
					num1[index1][2]++
					num2[index2][2]--
				}
			}
		}
		return
	}
	dfs(0, len(num2))
	if minCount == math.MaxInt {
		return 0
	}
	return minCount
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
