package leetcode

import (
	"sort"
)

//执行用时：4 ms, 在所有 Go 提交中击败了93.48%的用户
//内存消耗：3.2 MB, 在所有 Go 提交中击败了82.61%的用户
func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var sum int
	for index := range nums {
		sum += nums[index]
	}
	sum >>= 1
	var res []int
	for index := range nums {
		sum -= nums[index]
		res = append(res, nums[index])
		if sum < 0 {
			break
		}
	}
	return res
}
