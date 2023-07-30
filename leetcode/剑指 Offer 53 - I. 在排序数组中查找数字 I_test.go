package leetcode

import (
	"fmt"
	"testing"
)

func search(nums []int, target int) int {
	l := len(nums) - 1
	index := getIndex(nums, target, 0, l/2, l)
	if index == -1 {
		return 0
	}
	var count int
	for i := index; i <= l; i++ {
		if nums[i] == target {
			count++
		} else {
			break
		}
	}
	for i := index - 1; i >= 0; i-- {
		if nums[i] == target {
			count++
		} else {
			break
		}
	}
	return count
}

// 返回target所在下标，没有则返回-1
func getIndex(nums []int, target int, left, mid, right int) int {
	if left > right {
		return -1
	}
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		right = mid - 1
	} else {
		left = mid + 1
	}
	// 10^5	不会出现越界问题，所以无需考虑 (right+left)/2 是否越界，也可以使用 (right-left)/2 + left 防止越界
	return getIndex(nums, target, left, (right+left)/2, right)
}

func TestGetIndex(t *testing.T) {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
}
