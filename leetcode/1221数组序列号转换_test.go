package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"sort"
	"testing"
)

func arrayRankTransform(arr []int) []int {
	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] <= arr1[j]
	})
	m := make(map[int]int)
	for index := range arr1 {
		if _, ok := m[arr1[index]]; !ok {
			m[arr1[index]] = len(m) + 1
		}
	}
	for index := range arr {
		arr1[index] = m[arr[index]]
	}
	return arr1
}

//func arrayRankTransform(arr []int) []int {
//	min, max := math.MaxInt, math.MinInt
//	for _, num := range arr {
//		if num < min {
//			min = num
//		}
//		if num > max {
//			max = num
//		}
//	}
//	count := make([]int, max-min+1)
//	for _, num := range arr {
//		count[num-min] = 1
//	}
//
//	preSum := make([]int, len(count)+1)
//	for i := 1; i < len(preSum); i++ {
//		preSum[i] = preSum[i-1] + count[i-1]
//	}
//	ans := make([]int, len(arr))
//	index := 0
//	for _, num := range arr {
//		ans[index] = preSum[num-min] + 1
//		index++
//	}
//	return ans
//}

// 执行耗时:72 ms,击败了29.17% 的Go用户
// 内存消耗:12.3 MB,击败了52.08% 的Go用户
func TestArrayRankTransform(t *testing.T) {
	collections := []struct {
		input  []int
		output []int
	}{
		{
			[]int{40, 10, 20, 30},
			[]int{4, 1, 2, 3},
		},
		{
			[]int{100, 100, 100},
			[]int{1, 1, 1},
		},
		{
			[]int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			[]int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
		{
			[]int{-1000000000, 10, 20, 30, 1000000000},
			[]int{1, 2, 3, 4, 5},
		},
	}

	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, arrayRankTransform(collections[index].input))
	}
}
