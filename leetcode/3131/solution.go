package leetcode

// 执行耗时:5 ms,击败了31.00% 的Go用户
// 内存消耗:2.8 MB,击败了99.00% 的Go用户
func addedInteger(nums1 []int, nums2 []int) int {
	var max1, max2 int = nums1[0], nums2[0]
	for i := 1; i < len(nums1); i++ {
		max1 = max(max1, nums1[i])
	}
	for i := 1; i < len(nums2); i++ {
		max2 = max(max2, nums2[i])
	}
	return max2 - max1
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
