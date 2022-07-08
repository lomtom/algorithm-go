package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

func minCostToMoveChips(position []int) int {
	var ints [2]int
	for index := range position {
		ints[position[index]%2]++
	}
	if ints[0] > ints[1] {
		return ints[1]
	}
	return ints[0]
}

func TestMinCostToMoveChips(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{1, 2, 3},
			1,
		},
		{
			[]int{2, 2, 2, 3, 3},
			2,
		},
		{[]int{1, 1000000000},
			1,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, minCostToMoveChips(collections[index].input))
	}
}
