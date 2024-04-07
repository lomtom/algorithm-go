package leetcode

// 执行耗时:38 ms,击败了92.07% 的Go用户
// 内存消耗:6.7 MB,击败了75.46% 的Go用户
func canJump(nums []int) bool {
	var max = 0
	for index := 0; index < len(nums)-1; index++ {
		// 最远达到index处，并且无法再次跳跃
		if max == index && nums[index] == 0 {
			return false
		}
		if index+nums[index] > max {
			max = index + nums[index]
		}
	}
	return max >= len(nums)-1
}
