package leetcode

import "sort"

// 执行耗时:8 ms,击败了35.71% 的Go用户
// 内存消耗:2.8 MB,击败了78.57% 的Go用户
func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	sort.Ints(nums)
	var neg, pos, zero int
	var res int = 1
	var negMax int = -10
	for _, num := range nums {
		if num < 0 {
			res *= num
			neg++
			if num > negMax {
				negMax = num
			}
		} else if num > 0 {
			res *= num
			pos++
		} else {
			zero++
		}
	}
	if pos == 0 && neg <= 1 && zero != 0 {
		return 0
	}
	if neg%2 == 1 {
		return int64(res / negMax)
	}
	return int64(res)
}
