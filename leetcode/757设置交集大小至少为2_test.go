package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"sort"
	"testing"
)

func intersectionSizeTwo(intervals [][]int) (ans int) {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[1] < b[1] || a[1] == b[1] && a[0] > b[0]
	})
	a, b := -1, -1
	for _, interval := range intervals {
		if interval[0] > b {
			ans += 2
			a = interval[1] - 1
			b = interval[1]
		} else if interval[0] > a {
			a = b
			b = interval[1]
			ans++
		}
	}
	return
}

func TestIntersectionSizeTwo(t *testing.T) {
	collections := []struct {
		input  [][]int
		output int
	}{
		{
			[][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			3,
		},
		{
			[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			5,
		},
		{
			[][]int{{1, 2}},
			2,
		},
		{
			[][]int{{1, 3}, {3, 7}, {5, 7}, {7, 8}},
			5,
		},
	}

	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, intersectionSizeTwo(collections[index].input))
	}
}
