package leetcode

import (
	"sort"
)

// 执行耗时:44 ms,击败了10.67% 的Go用户
// 内存消耗:9.8 MB,击败了69.63% 的Go用户
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var left, right = i + 1, len(nums) - 1
		for left < right {
			for left < right && left > i+1 && nums[left] == nums[left-1] {
				left++
			}
			for left < right && right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
			}
			if left >= right {
				continue
			}
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return res
}
