package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"sort"
	"testing"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var min int = int(^uint(0) >> 1)
	for i := 1; i < len(arr); i++ {
		temp := arr[i] - arr[i-1]
		if temp < min {
			min = temp
		}
	}
	res := make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}

func TestMinimumAbsDifference(t *testing.T) {
	collections := []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{4, 2, 1, 3},
			output: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			[]int{1, 3, 6, 10, 15},
			[][]int{{1, 3}},
		},
	}

	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(minimumAbsDifference(collections[index].input), collections[index].output)
	}
}
