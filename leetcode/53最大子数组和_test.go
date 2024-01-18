package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:87 ms,击败了42.34% 的Go用户
// 内存消耗:8.3 MB,击败了45.34% 的Go用户
func maxSubArray(nums []int) int {
	maxFunc := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	maxNumber := nums[0]
	for index := 1; index < len(nums); index++ {
		nums[index] = maxFunc(nums[index], nums[index]+nums[index-1])
		if nums[index] > maxNumber {
			maxNumber = nums[index]
		}
	}
	return maxNumber
}

func TestMaxSubArray(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{[]int{1},
			1,
		},
		{[]int{5, 4, -1, 7, 8},
			23,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, maxSubArray(collections[index].input))
	}
}
