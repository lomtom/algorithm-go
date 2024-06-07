package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.3 MB,击败了36.36% 的Go用户
func maxOperations(nums []int) (result int) {
	sum := nums[0] + nums[1]
	for i := 0; i+1 < len(nums); i = i + 2 {
		if sum != nums[i]+nums[i+1] {
			break
		}
		result++
	}
	return
}
