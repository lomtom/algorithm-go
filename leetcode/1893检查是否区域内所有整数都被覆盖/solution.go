package leetcode

import (
	"fmt"
	"testing"
)

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.4 MB,击败了100.00% 的Go用户
func isCovered(ranges [][]int, left int, right int) bool {
	area := make([]bool, right-left+1)
	for index := range ranges {
		for i := ranges[index][0]; i <= ranges[index][1]; i++ {
			if i >= left && i <= right {
				area[i-left] = true
			}
		}
	}
	for index := range area {
		if area[index] == false {
			return false
		}
	}
	return true
}

func TestIsCovered(t *testing.T) {
	fmt.Println(isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
	fmt.Println(isCovered([][]int{{1, 10}, {10, 20}}, 21, 21))
}
