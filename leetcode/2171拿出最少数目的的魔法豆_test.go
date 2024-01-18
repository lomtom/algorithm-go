package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"sort"
	"testing"
)

// 执行耗时:189 ms,击败了90.00% 的Go用户
// 内存消耗:8.2 MB,击败了95.00% 的Go用户
func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	var total int64
	for index := range beans {
		total += int64(beans[index])
	}
	min := func(i, j int64) int64 {
		if i > j {
			return j
		}
		return i
	}
	var res = total
	for index := range beans {
		res = min(res, total-int64(beans[index])*int64(len(beans)-index))
	}
	return res
}

func TestMinimumRemoval(t *testing.T) {
	collections := []struct {
		input  []int
		output int64
	}{
		{
			[]int{4, 1, 6, 5},
			4,
		},
		{[]int{2, 10, 3, 2},
			7,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, minimumRemoval(collections[index].input))
	}
}
