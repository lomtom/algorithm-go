package leetcode

import (
	"sort"
)

// 执行耗时:49 ms,击败了69.57% 的Go用户
// 内存消耗:6.65 MB,击败了100.00% 的Go用户
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sort.Ints(aliceSizes)
	sort.Ints(bobSizes)
	sum := 0
	for i := 0; i < len(aliceSizes); i++ {
		sum += aliceSizes[i]
	}
	for i := 0; i < len(bobSizes); i++ {
		sum -= bobSizes[i]
	}
	sub := sum / 2
	for i, j := 0, 0; i < len(aliceSizes) && j < len(bobSizes); {
		temp := aliceSizes[i] - bobSizes[j]
		if temp == sub {
			return []int{aliceSizes[i], bobSizes[j]}
		} else if temp > sub {
			j++
		} else {
			i++
		}
	}
	return []int{aliceSizes[len(aliceSizes)-1], bobSizes[len(bobSizes)-1]}
}
