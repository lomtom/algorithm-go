package leetcode

import "sort"

// 执行耗时:13 ms,击败了98.91% 的Go用户
// 内存消耗:6.34 MB,击败了65.57% 的Go用户
func minimumPushes(word string) int {
	var pairs [26]int
	for index := range word {
		pairs[word[index]-'a']++
	}
	sort.Slice(pairs[:], func(i, j int) bool {
		return pairs[i] > pairs[j]
	})
	var level int = 0
	var result int
	for index := range pairs {
		if index%8 == 0 {
			level++
		}
		result += level * pairs[index]
	}
	return result
}
