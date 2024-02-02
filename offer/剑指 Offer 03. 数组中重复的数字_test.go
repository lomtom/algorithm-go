package offer

import "sort"

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return 0
}

func findRepeatNumber1(nums []int) int {
	sort.Ints(nums)
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		if num == nums[i] {
			break
		}
		num = nums[i]
	}
	return num
}
