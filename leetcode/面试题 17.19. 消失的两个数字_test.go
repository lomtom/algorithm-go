package leetcode

import (
	"fmt"
	"testing"
)

func missingTwo(nums []int) []int {
	var n = len(nums) + 2
	var cur = n * (1 + n) / 2
	for index := range nums {
		cur -= nums[index]
	}
	var sum, t = cur, cur / 2
	cur = t * (1 + t) / 2
	for index := range nums {
		if nums[index] <= t {
			cur -= nums[index]
		}
	}
	return []int{cur, sum - cur}
}

func TestMissingTwo(t *testing.T) {
	fmt.Println(missingTwo([]int{}))
	fmt.Println(missingTwo([]int{2, 3}))
	fmt.Println(missingTwo([]int{1}))
}
