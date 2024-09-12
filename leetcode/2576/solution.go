package leetcode

import "sort"

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, n/2
	var count int = l
	for j := r; count < r && j < n; {
		if nums[count]*2 <= nums[j] {
			count++
			j++
		} else {
			j++
		}
	}
	return count * 2
}
