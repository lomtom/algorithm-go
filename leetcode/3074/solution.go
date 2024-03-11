package leetcode

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	var result int
	var appleTotal int
	for index := range apple {
		appleTotal += apple[index]
	}
	sort.Ints(capacity)
	var index int = len(capacity) - 1
	for appleTotal > 0 {
		appleTotal -= capacity[index]
		result++
		index--
	}
	return result
}
