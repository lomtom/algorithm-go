package leetcode

import (
	"math"
	"sort"
)

func findLongestChain(pairs [][]int) (res int) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	cur := math.MinInt32
	for index := range pairs {
		if pairs[index][0] > cur {
			cur = pairs[index][1]
			res++
		}
	}
	return
}
