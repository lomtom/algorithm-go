package leetcode

import "math"

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了67.86% 的Go用户
func constructRectangle(area int) []int {
	for i := int(math.Sqrt(float64(area))); i >= 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{}
}
