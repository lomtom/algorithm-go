package leetcode

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	var count int
	var result int64
	var index int = len(happiness) - 1
	for k > 0 {
		temp := happiness[index] - count
		if temp > 0 {
			result += int64(temp)
		}
		count++
		index--
		k--
	}
	return result
}
