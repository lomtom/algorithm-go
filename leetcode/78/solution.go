package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.3 MB,击败了45.65% 的Go用户
func subsets(nums []int) (result [][]int) {
	var dfs func(index int, res []int)
	dfs = func(index int, res []int) {
		var ans = make([]int, len(res))
		copy(ans, res)
		result = append(result, ans)
		for i := index; i < len(nums); i++ {
			res = append(res, nums[i])
			dfs(i+1, res)
			res = res[:len(res)-1]
		}
	}
	dfs(0, []int{})
	return
}
