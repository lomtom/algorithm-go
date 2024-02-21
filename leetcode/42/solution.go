package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// 执行耗时:8 ms,击败了89.64% 的Go用户
// 内存消耗:5.9 MB,击败了24.68% 的Go用户
func trap(height []int) int {
	var maxFunc = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	var minFunc = func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	var l = len(height)
	var maxL = make([]int, l)
	var maxR = make([]int, l)
	for i := 0; i < l; i++ {
		maxL[i] = height[i]
		maxR[l-1-i] = height[l-1-i]
		if i != 0 {
			maxL[i] = maxFunc(maxL[i], maxL[i-1])
			maxR[l-1-i] = maxFunc(maxR[l-1-i], maxR[l-i])
		}
	}
	var total int
	for i := 0; i < l; i++ {
		total += minFunc(maxL[i], maxR[i]) - height[i]
	}
	return total
}

func TestTrap(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{[]int{4, 2, 0, 3, 2, 5},
			9,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, trap(collections[index].input))
	}
}
