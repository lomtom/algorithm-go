package leetcode

import "sort"

// 执行耗时:45 ms,击败了61.14% 的Go用户
// 内存消耗:9.3 MB,击败了13.36% 的Go用户
func accountsMerge(accounts [][]string) [][]string {
	var emailToIndex = make(map[string]int)
	var emailToName = make(map[string]string)
	for index := range accounts {
		for i := 1; i < len(accounts[index]); i++ {
			if _, has := emailToIndex[accounts[index][i]]; !has {
				emailToIndex[accounts[index][i]] = len(emailToIndex)
				emailToName[accounts[index][i]] = accounts[index][0]
			}
		}
	}
	parent := make([]int, len(emailToIndex))
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union func(x, y int)
	union = func(x, y int) {
		parent[find(x)] = find(y)
	}
	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}
	var indexToEmails = make(map[int][]string)
	for email, index := range emailToIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}
	var res [][]string
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		name := emailToName[emails[0]]
		res = append(res, append([]string{name}, emails...))
	}
	return res
}
