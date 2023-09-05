package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"math"
	"testing"
)

func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}

	level := 0
	maxLevel, maxValue := 1, math.MinInt
	Func := func() {
		level++
		l := len(queue)
		value := 0
		for i := 0; i < l; i++ {
			node := queue[i]
			if node != nil {
				value += node.Val
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}
		if value > maxValue {
			maxValue = value
			maxLevel = level
		}
		queue = queue[l:]
	}
	for len(queue) != 0 {
		Func()
	}
	return maxLevel
}

// 执行用时：92 ms, 在所有 Go 提交中击败了94.00%的用户
// 内存消耗：8.5 MB, 在所有 Go 提交中击败了6.00%的用户
func TestMaxLevelSum(t *testing.T) {
	collections := []struct {
		input  []int
		output int
	}{
		{
			[]int{1, 7, 0, 7, -8},
			2,
		},
	}

	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, maxLevelSum(construct(collections[index].input)))
	}
}
