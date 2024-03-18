package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

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

func TestNumArray1(t *testing.T) {
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
		obj := Constructor1(collections[index].input.nums)
		for i := range collections[index].indexs {
			s.Equal(collections[index].output[i], obj.SumRange(collections[index].input.indexs[i][0], collections[index].input.indexs[i][1]))
		}

	}
}
