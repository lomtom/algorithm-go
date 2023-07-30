package leetcode

import (
	"fmt"
	"testing"
)

func missingNumber(nums []int) int {
	right := len(nums) - 1
	left := 0
	mid := right / 2
	for left <= right {
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (right-left)/2 + left
	}
	return mid
}

func TestMissingNumber(t *testing.T) {
	fmt.Println(missingNumber([]int{0, 1, 3}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	fmt.Println(missingNumber([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 6}))
	fmt.Println(missingNumber([]int{0}))
}
