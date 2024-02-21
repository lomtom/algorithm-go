package leetcode

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：3.4 MB, 在所有 Go 提交中击败了12.82%的用户
func nextGreaterElement(nums1 []int, nums2 []int) (ans []int) {
	var stack []int
	m := make(map[int]int, 0)
	for index := len(nums2) - 1; index >= 0; index-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[index] {
			stack = stack[:len(stack)-1]
		}
		var res = -1
		if len(stack) > 0 {
			res = stack[len(stack)-1]
		}
		m[nums2[index]] = res
		stack = append(stack, nums2[index])
	}
	for index := range nums1 {
		ans = append(ans, m[nums1[index]])
	}
	return
}
