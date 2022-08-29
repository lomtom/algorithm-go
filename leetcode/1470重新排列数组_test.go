package leetcode

//执行用时：4 ms, 在所有 Go 提交中击败了86.57%的用户
//内存消耗：3.7 MB, 在所有 Go 提交中击败了19.40%的用户
func shuffle(nums []int, n int) (ans []int) {
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i], nums[i+n])
	}
	return
}
