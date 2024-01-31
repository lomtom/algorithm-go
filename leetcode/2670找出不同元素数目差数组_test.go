package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:22 ms,击败了26.09% 的Go用户
// 内存消耗:5.5 MB,击败了56.52% 的Go用户
func distinctDifferenceArray(nums []int) []int {
	preM := make(map[int]int)
	sufM := make(map[int]int)
	preCount := 0
	sufCount := 0
	// 预处理后缀
	for index := range nums {
		if count, ok := sufM[nums[index]]; ok {
			sufM[nums[index]] = count + 1
		} else {
			sufM[nums[index]] = 1
			sufCount++
		}
	}
	result := make([]int, len(nums))
	for index := range nums {
		// 处理前缀
		if count, ok := preM[nums[index]]; ok {
			preM[nums[index]] = count + 1
		} else {
			preM[nums[index]] = 1
			preCount++
		}
		// 处理后缀
		sufM[nums[index]] = sufM[nums[index]] - 1
		if sufM[nums[index]] == 0 {
			sufCount--
		}
		// 计算结果
		result[index] = preCount - sufCount
	}
	return result
}

func TestDistinctDifferenceArray(t *testing.T) {
	collections := []struct {
		input  []int
		output []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{-3, -1, 1, 3, 5},
		},
		{[]int{3, 2, 3, 4, 2},
			[]int{-2, -1, 0, 2, 3},
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, distinctDifferenceArray(collections[index].input))
	}
}
