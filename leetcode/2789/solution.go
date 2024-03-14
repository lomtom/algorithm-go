package leeetcode

// 执行耗时:117 ms,击败了10.71% 的Go用户
// 内存消耗:9.3 MB,击败了10.71% 的Go用户
func maxArrayValue(nums []int) int64 {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] >= nums[i-1] {
			nums[i-1] += nums[i]
			nums[i] = 0
		}
	}
	var result = nums[len(nums)-1]
	for i := range nums {
		if nums[i] > result {
			result = nums[i]
		}
	}
	return int64(result)
}
