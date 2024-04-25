package leetcode

// 执行耗时:2 ms,击败了59.02% 的Go用户
// 内存消耗:2.0 MB,击败了13.74% 的Go用户
func combinationSum3(k int, n int) (result [][]int) {
	var dfs func(index, sub int, res []int)
	dfs = func(index, sub int, res []int) {
		if sub == 0 && k == 0 {
			result = append(result, append([]int{}, res...))
			return
		}
		for i := index; i <= 9; i++ {
			if k > 0 && sub-i >= 0 {
				res = append(res, i)
				k--
				dfs(i+1, sub-i, res)
				res = res[:len(res)-1]
				k++
			}
		}
	}
	dfs(1, n, []int{})
	return result
}
