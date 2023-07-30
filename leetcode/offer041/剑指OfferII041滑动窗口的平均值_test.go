package offer041

import (
	"fmt"
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

type MovingAverage struct {
	// 保存数量
	num int
	// 总值
	sum  int
	nums []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		0,
		0,
		make([]int, size),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	index := this.num % len(this.nums)
	this.sum -= this.nums[index]
	this.sum += val
	this.nums[index] = val
	this.num++
	if this.num < len(this.nums) {
		return float64(this.sum) / float64(this.num)
	}
	return float64(this.sum) / float64(len(this.nums))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func TestMovingAverage(t *testing.T) {
	type input struct {
		movingAverage int
		nexts         []int
	}
	collections := []struct {
		input
		output []float64
	}{
		{
			input{
				3,
				[]int{1, 10, 3, 5},
			},
			[]float64{1, 5.5, 4.66667, 6},
		},
		{
			input{
				1,
				[]int{1, 2, 3, 4, 5},
			},
			[]float64{1, 2, 3, 4, 5},
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		obj := Constructor(collections[index].input.movingAverage)
		for index2, value := range collections[index].input.nexts {
			s.Equal(fmt.Sprintf("%.5f", collections[index].output[index2]), fmt.Sprintf("%.5f", obj.Next(value)))
		}
	}
}
