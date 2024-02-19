package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了75.77% 的Go用户
func removeElement(nums []int, val int) int {
	var curIndex = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[curIndex] = nums[i]
			curIndex++
		}
	}
	return curIndex
}

func TestRemoveElement(t *testing.T) {
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
				[]int{3, 2, 2, 3},
				3,
			},
			2,
		},
		{
			input{
				[]int{0, 1, 2, 2, 3, 0, 4, 2},
				2,
			},
			5,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, removeElement(collections[index].input.nums, collections[index].input.target))
	}
}
