package leetcode

import "sort"

// 执行耗时:366 ms,击败了42.86% 的Go用户
// 内存消耗:8.4 MB,击败了14.29% 的Go用户
func maxmiumScore(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	var result = 0
	var odd, even = -1, -1
	for index := 0; index < cnt; index++ {
		result += cards[index]
		if cards[index]%2 == 0 {
			even = index
		} else {
			odd = index
		}
	}
	if result%2 == 0 {
		return result
	}
	for index := cnt; index < len(cards); index++ {
		if cards[index]%2 == 1 && even != -1 {
			result += cards[index] - cards[even]
			return result
		} else if cards[index]%2 == 0 && odd != -1 {
			result += cards[index] - cards[odd]
			return result
		}
	}
	return 0
}
