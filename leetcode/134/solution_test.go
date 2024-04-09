package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func Test(t *testing.T) {
	type input struct {
		gas  []int
		cost []int
	}
	collections := []struct {
		input  input
		output int
	}{
		{
			input{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
			3,
		},
		{
			input{[]int{5, 4, 3, 2, 1}, []int{2, 1, 5, 4, 3}},
			0,
		},
		{
			input{[]int{2, 3, 4}, []int{3, 4, 5}},
			-1,
		},
		{ //-2,3,-4,5,5
			input{[]int{6, 6, 6, 6, 6}, []int{8, 3, 10, 1, 1}},
			3,
		},
		{ // 2,1,0,0,0,0
			input{[]int{2, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 0, 0}},
			0,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, canCompleteCircuit(collections[index].input.gas, collections[index].input.cost))
	}
}
