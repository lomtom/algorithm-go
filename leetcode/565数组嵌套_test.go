package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func arrayNesting(nums []int) (ans int) {
	vis := make([]bool, len(nums))
	for index := range nums {
		res := 0
		for !vis[index] {
			res++
			vis[index] = true
			index = nums[index]
		}
		if res > ans {
			ans = res
		}
	}
	return
}

func TestArrayNesting(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{5, 4, 0, 3, 1, 6, 2},
			4,
		},
		{
			[]int{0},
			1,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, arrayNesting(collections[index].input))
	}
}
