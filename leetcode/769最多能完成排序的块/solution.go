package leetcode

import (
	"fmt"
	"testing"
)

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.8 MB,击败了75.41% 的Go用户
func maxChunksToSorted(arr []int) (ans int) {
	v := 0
	for index, value := range arr {
		if value > v {
			v = value
		}
		if index == v {
			ans++
		}
	}
	return
}

func TestMaxChunksToSorted(t *testing.T) {
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
	fmt.Println(maxChunksToSorted([]int{2, 0, 1}))
}
