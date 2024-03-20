package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.18 MB,击败了6.45% 的Go用户
func dominantIndex(nums []int) int {
	max := 0
	for index := range nums {
		if nums[index] > nums[max] {
			max = index
		}
	}
	for index := range nums {
		if 2*nums[index] > nums[max] && max != index {
			return -1
		}
	}
	return max
}
