package leetcode

import "sort"

// 执行耗时:24 ms,击败了18.15% 的Go用户
// 内存消耗:7.5 MB,击败了72.38% 的Go用户
func groupAnagrams(strs []string) [][]string {
	var m = make(map[string][]string)
	for _, str := range strs {
		var bytes = []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		m[string(bytes)] = append(m[string(bytes)], str)
	}
	var ans [][]string
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
