package _03

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

type NumArray struct {
	nums []int
}

func lowBit(index int) int {
	return index & -index
}

func sum(ints []int, index int) (ans int) {
	for index > 0 {
		ans += ints[index]
		index -= lowBit(index)
	}
	return ans
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	ints := make([]int, l+1)
	for index := 1; index <= l; index++ {
		index1 := index
		for index1 <= l {
			ints[index1] += nums[index-1]
			index1 += lowBit(index1)
		}
	}
	return NumArray{
		ints,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return sum(this.nums, right+1) - sum(this.nums, left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func TestNumArray(t *testing.T) {
	type input struct {
		nums   []int
		indexs [][2]int
	}
	{
	}
	collections := []struct {
		input
		output []int
	}{
		{
			input{
				nums:   []int{-2, 0, 3, -5, 2, -1},
				indexs: [][2]int{{0, 2}, {2, 5}, {0, 5}},
			},
			[]int{1, -1, -3},
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		obj := Constructor(collections[index].input.nums)
		for i := range collections[index].indexs {
			s.Equal(collections[index].output[i], obj.SumRange(collections[index].input.indexs[i][0], collections[index].input.indexs[i][1]))
		}

	}
}
