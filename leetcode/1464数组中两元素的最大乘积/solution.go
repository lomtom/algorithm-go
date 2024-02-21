package leetcode

import (
	"fmt"
	"math"
	"testing"
)

//执行耗时:4 ms,击败了62.03% 的Go用户
//内存消耗:2.8 MB,击败了79.75% 的Go用户
//func maxProduct(nums []int) int {
//	sort.Ints(nums)
//	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
//}

// 执行耗时:4 ms,击败了62.03% 的Go用户
// 内存消耗:2.8 MB,击败了100.00% 的Go用户
func maxProduct(nums []int) int {
	var index1 = 0
	var max1 = math.MinInt
	for index := 0; index < len(nums); index++ {
		if nums[index] > max1 {
			index1 = index
			max1 = nums[index]
		}
	}
	var max2 = math.MinInt
	for index := 0; index < len(nums); index++ {
		if nums[index] > max2 && index != index1 {
			max2 = nums[index]
		}
	}
	return (max1 - 1) * (max2 - 1)
}

func Test(t *testing.T) {
	fmt.Println(maxProduct([]int{10, 2, 5, 2}))
}
