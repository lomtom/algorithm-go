package leetcode

import (
	"fmt"
	"testing"
)

// 方法一：暴力解法
//执行用时：320 ms, 在所有 Go 提交中击败了5.01%的用户
//内存消耗：6.3 MB, 在所有 Go 提交中击败了100.00%的用户
//func nextGreaterElements(nums []int) []int {
//	ans := make([]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		index := (i + 1) % len(nums)
//		for nums[index] <= nums[i] && index != i {
//			index = (index + 1) % len(nums)
//		}
//		if index == i {
//			ans[i] = -1
//		} else {
//			ans[i] = nums[index]
//		}
//	}
//	return ans
//}

//执行用时：20 ms, 在所有 Go 提交中击败了93.87%的用户
//内存消耗：6.5 MB, 在所有 Go 提交中击败了92.20%的用户
func nextGreaterElements(nums []int) []int {
	l := len(nums)
	stack := make([]int, 0)
	ans := make([]int, l)
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i < 2*l; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%l] {
			ans[stack[len(stack)-1]] = nums[i%l]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%l)
	}
	return ans
}

func TestNextGreaterElements(t *testing.T) {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
