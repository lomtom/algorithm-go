package leetcode

// 执行耗时:4 ms,击败了40.63% 的Go用户
// 内存消耗:2.7 MB,击败了100.00% 的Go用户
func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	var ans int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				ans++
			}
		}
	}
	return ans
}
