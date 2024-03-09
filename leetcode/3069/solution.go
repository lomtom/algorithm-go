package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:3 MB,击败了88.33% 的Go用户
func resultArray(nums []int) []int {
	a := []int{nums[0]}
	b := []int{nums[1]}
	for i := 2; i < len(nums); i++ {
		if a[len(a)-1] > b[len(b)-1] {
			a = append(a, nums[i])
		} else {
			b = append(b, nums[i])
		}
	}
	return append(a, b...)
}
