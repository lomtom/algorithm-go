package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 2,7,9,3,1
// 1
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了70.13% 的Go用户
func rob(nums []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for index := range nums {
		if index == 1 {
			nums[index] = max(nums[index], nums[index-1])
		} else if index > 1 {
			nums[index] = max(nums[index-2]+nums[index], nums[index-1])
		}
	}
	return nums[len(nums)-1]
}

func TestRob(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{2, 7, 9, 3, 1},
			12,
		},
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{[]int{2},
			2,
		},
		{[]int{7, 2},
			7,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, rob(collections[index].input))
	}
}
