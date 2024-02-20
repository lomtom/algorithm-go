package leetcode

import (
	"testing"
)

// [0,1,0,3,12] -> [1,3,12,0,0]
// 执行耗时:21 ms,击败了20.38% 的Go用户
// 内存消耗:6.5 MB,击败了73.43% 的Go用户
func moveZeroes(nums []int) {
	var left, right int = 0, 0
	for ; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}

func TestMoveZeroes(t *testing.T) {
	var nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}
