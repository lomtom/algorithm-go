package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)
	for index := range groupSizes {
		array := m[groupSizes[index]]
		array = append(array, index)
		m[groupSizes[index]] = array
	}
	ans := make([][]int, 0)
	for key, value := range m {
		res := make([]int, 0)
		for i := range value {
			res = append(res, value[i])
			if (i+1)%key == 0 || i+1 == len(value) {
				ans = append(ans, res)
				res = make([]int, 0)
			}
		}
	}
	return ans
}

// 执行耗时:8 ms,击败了40.91% 的Go用户
// 内存消耗:5.6 MB,击败了18.18% 的Go用户
func TestGroupThePeople(t *testing.T) {
	collections := []struct {
		input  []int
		output [][]int
	}{
		{
			[]int{3, 3, 3, 3, 3, 1, 3},
			[][]int{
				{5}, {0, 1, 2}, {3, 4, 6},
			},
		},
		{
			[]int{2, 1, 3, 3, 3, 2},
			[][]int{{1}, {0, 5}, {2, 3, 4}},
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, groupThePeople(collections[index].input))
	}
}
