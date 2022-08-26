package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}

//执行耗时:4 ms,击败了62.03% 的Go用户
//内存消耗:2.8 MB,击败了79.75% 的Go用户
func Test(t *testing.T) {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}
