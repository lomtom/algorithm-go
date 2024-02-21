package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"sort"
	"testing"
)

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var index1, index2 = 0, 0
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] == nums2[index2] {
			return nums1[index1]
		} else if nums1[index1] < nums2[index2] {
			index1++
		} else {
			index2++
		}
	}
	if nums1[0] < nums2[0] {
		return nums1[0]*10 + nums2[0]
	}
	return nums2[0]*10 + nums1[0]
}

type input struct {
	nums1 []int
	nums2 []int
}

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了41.38% 的Go用户
func TestMinNumber(t *testing.T) {
	collections := []struct {
		input  input
		output int
	}{
		{
			input{
				[]int{4, 1, 3},
				[]int{5, 7},
			},
			15,
		},
		{
			input{
				[]int{3, 5, 2, 6},
				[]int{3, 1, 7},
			},
			3,
		},
		{
			input{
				[]int{7, 5, 6},
				[]int{1, 4},
			},
			15,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, minNumber(collections[index].input.nums1, collections[index].input.nums2))
	}
}
