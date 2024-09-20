package leetcode

const l = 1e5 + 10

// 执行耗时:42 ms,击败了85.79% 的Go用户
// 内存消耗:7.7 MB,击败了73.78% 的Go用户
func firstMissingPositive(nums []int) int {
	var flags = [l]bool{}
	for _, num := range nums {
		if num >= 0 && num <= l {
			flags[num] = true
		}
	}
	for i := 1; i < l; i++ {
		if !flags[i] {
			return i
		}
	}
	return 0
}
