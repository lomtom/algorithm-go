package leetcode

// 执行耗时:11 ms,击败了48.61% 的Go用户
// 内存消耗:5.7 MB,击败了100.00% 的Go用户
func maximumCount(nums []int) int {
	number1 := count(nums, 0)
	number2 := len(nums) - count(nums, 1)
	if number1 > number2 {
		return number1
	}
	return number2
}

func count(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
