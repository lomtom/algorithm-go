package leetcode

func minCostToMoveChips(position []int) int {
	var ints [2]int
	for index := range position {
		ints[position[index]%2]++
	}
	if ints[0] > ints[1] {
		return ints[1]
	}
	return ints[0]
}
