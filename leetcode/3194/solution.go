package leetcode

import (
	"math"
	"sort"
)

// 执行耗时:4 ms,击败了38.64% 的Go用户
// 内存消耗:4 MB,击败了6.82% 的Go用户
func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	l := len(nums) / 2
	var ans float64 = math.MaxFloat64
	for i := 0; i < l; i++ {
		ans = math.Min(ans, float64(nums[len(nums)-1-i]+nums[i])/2)
	}
	return ans
}
