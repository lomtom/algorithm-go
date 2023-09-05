package _07

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

type NumArray struct {
	nums []int
	tree []int
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	tree := make([]int, l+1)
	for i := 1; i <= l; i++ {
		index := i
		for index <= l {
			tree[index] += nums[i-1]
			index += lowBit(index)
		}
	}
	return NumArray{
		nums: nums,
		tree: tree,
	}
}

func (this *NumArray) Update(index int, val int) {
	subVal := val - this.nums[index]
	this.nums[index] = val
	index++
	for index < len(this.tree) {
		this.tree[index] += subVal
		index += lowBit(index)
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return sum(right+1, this.tree) - sum(left, this.tree)
}

func lowBit(index int) int {
	return index & -index
}

func sum(index int, a []int) (ans int) {
	for index > 0 {
		ans += a[index]
		index -= lowBit(index)
	}
	return
}

func TestNumArray(t *testing.T) {
	type input struct {
		nums   []int
		indexs [][3]int
	}
	collections := []struct {
		input
		output []int
	}{
		{
			input{
				nums:   []int{1, 3, 5},
				indexs: [][3]int{{1, 0, 2}, {2, 1, 2}, {1, 0, 2}},
			},
			[]int{9, 0, 8},
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		obj := Constructor(collections[index].input.nums)
		for i := range collections[index].indexs {
			if collections[index].input.indexs[i][0] == 1 {
				s.Equal(collections[index].output[i], obj.SumRange(collections[index].input.indexs[i][1], collections[index].input.indexs[i][2]))
			} else {
				obj.Update(collections[index].input.indexs[i][1], collections[index].input.indexs[i][2])
			}
		}

	}
}
