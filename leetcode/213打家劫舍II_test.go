package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了70.13% 的Go用户
func robII(nums []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	var first1, second1 = nums[0], max(nums[0], nums[1])
	var first2, second2 = nums[1], max(nums[1], nums[2])
	for index := range nums[2:] {
		if index != 0 {
			first2, second2 = second2, max(nums[index]+first2, second2)
		}
		if index != len(nums)-3 {
			first1, second1 = second1, max(nums[index]+first1, second1)
		}
	}
	return max(second1, second2)
}

func TestRobII(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{2, 3, 2},
			3,
		},
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{[]int{1, 2, 3},
			3,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, robII(collections[index].input))
	}
}
