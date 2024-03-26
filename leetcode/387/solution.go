package leetcode

// 执行耗时:3 ms,击败了94.12% 的Go用户
// 内存消耗:5.9 MB,击败了35.62% 的Go用户
func firstUniqChar(s string) int {
	m := [26]int{}
	for index := range s {
		m[s[index]-'a']++
	}
	for index := range s {
		if m[s[index]-'a'] == 1 {
			return index
		}
	}
	return -1
}

// 执行耗时:50 ms,击败了6.86% 的Go用户
// 内存消耗:6 MB,击败了12.42% 的Go用户
func firstUniqChar1(s string) int {
	m := make(map[byte]int)
	for index := range s {
		m[s[index]]++
	}
	for index := range s {
		if m[s[index]] == 1 {
			return index
		}
	}
	return -1
}
