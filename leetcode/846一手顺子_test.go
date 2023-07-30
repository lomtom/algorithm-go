package leetcode

import "sort"

//[1,2,3,3,4,4,5,6],4 true
//[3,2,1,2,3,4,3,4,5,9,10,11],3 true
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize > 0 {
		return false
	}
	sort.Ints(hand)
	cnt := map[int]int{}
	for _, num := range hand {
		cnt[num]++
	}
	for _, x := range hand {
		if cnt[x] == 0 {
			continue
		}
		for num := x; num < x+groupSize; num++ {
			if cnt[num] == 0 {
				return false
			}
			cnt[num]--
		}
	}
	return true
}
