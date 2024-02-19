package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:26 ms,击败了60.66% 的Go用户
// 内存消耗:6.5 MB,击败了50.75% 的Go用户
func search(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func TestSearch(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	collections := []struct {
		input  input
		output int
	}{
		{
			input{
				[]int{-1, 0, 3, 5, 9, 12},
				13,
			},
			-1,
		},
		{
			input{
				[]int{-1, 0, 3, 5, 9, 12},
				9,
			},
			4,
		},
		{
			input{
				[]int{-1, 0, 3, 5, 9, 12},
				2,
			},
			-1,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, search(collections[index].input.nums, collections[index].input.target))
	}
}
