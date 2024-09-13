package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了73.53% 的Go用户
func robII(nums []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	var first1, second1 = nums[0], max(nums[0], nums[1])
	var first2, second2 = nums[1], max(nums[1], nums[2])
	for index := 2; index < len(nums); index++ {
		// 不偷取第一家的情况
		if index != 2 {
			first2, second2 = second2, max(nums[index]+first2, second2)
		}
		// 不偷取最后一家的情况
		if index != len(nums)-1 {
			first1, second1 = second1, max(nums[index]+first1, second1)
		}
	}
	return max(second1, second2)
}
