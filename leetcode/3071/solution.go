package leetcode

import "math"

// 执行耗时:39 ms,击败了66.90% 的Go用户
// 内存消耗:6.8 MB,击败了81.69% 的Go用户
func minimumOperationsToWriteY(grid [][]int) int {
	var y [3]int
	var notY [3]int
	rowMid := len(grid) / 2
	colMid := len(grid[0]) / 2
	for i := range grid {
		for j := range grid[i] {
			if isY(i, j, rowMid, colMid) {
				y[grid[i][j]]++
			} else {
				notY[grid[i][j]]++
			}
		}
	}
	var res = math.MaxInt
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				res = min(res, getCount(y, notY, i, j))
			}
		}
	}
	return res
}
func isY(i, j, rowMid, colMid int) bool {
	if i <= rowMid {
		if colMid-i == j-colMid || i == j {
			return true
		}
		return false
	}
	if j == rowMid {
		return true
	}
	return false
}

func getCount(y [3]int, notY [3]int, yValue, notYValue int) int {
	var count int
	for i := 0; i < 3; i++ {
		if i != yValue {
			count += y[i]
		}
		if i != notYValue {
			count += notY[i]
		}
	}
	return count
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
