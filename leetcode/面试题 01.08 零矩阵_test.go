package leetcode

import "testing"

//执行耗时:12 ms,击败了69.55% 的Go用户
//内存消耗:6.1 MB,击败了88.64% 的Go用户
func setZeroes(matrix [][]int) {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 0 {
				m1[row] = true
				m2[col] = true
			}
		}
	}
	for row := range matrix {
		for col := range matrix[row] {
			if m1[row] || m2[col] {
				matrix[row][col] = 0
			}
		}
	}
}

func TestSetZeroes(t *testing.T) {
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
	setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
}
