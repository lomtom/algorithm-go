package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func Test(t *testing.T) {
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
