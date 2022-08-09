package leetcode

import "math"

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：1.9 MB, 在所有 Go 提交中击败了68.42%的用户
func minStartValue(nums []int) int {
	sum := 0
	min := math.MaxInt
	for index := range nums {
		sum += nums[index]
		if min > sum {
			min = sum
		}
	}
	if min <= 0 {
		return 1 - min
	} else {
		return 1
	}
}
