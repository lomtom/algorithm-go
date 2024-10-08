package leetcode

// 执行耗时:3 ms,击败了85.71% 的Go用户
// 内存消耗:3.9 MB,击败了71.59% 的Go用户
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			return []int{i, m[target-nums[i]]}
		}
		m[nums[i]] = i
	}
	return nil
}
