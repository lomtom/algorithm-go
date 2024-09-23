package leetcode

import "math"

// 执行耗时:30 ms,击败了11.90% 的Go用户
// 内存消耗:6.5 MB,击败了42.86% 的Go用户
func findShortestSubArray(nums []int) int {
	var maxC int
	var countMap = make(map[int]int)
	var firstMap = make(map[int]int)
	var lastMap = make(map[int]int)
	for index := range nums {
		countMap[nums[index]]++
		if countMap[nums[index]] > maxC {
			maxC = countMap[nums[index]]
		}
		if _, ok := firstMap[nums[index]]; !ok {
			firstMap[nums[index]] = index
		}
		lastMap[nums[index]] = index
	}
	var ans = math.MaxInt
	for index := range countMap {
		if countMap[index] == maxC {
			ans = min(ans, lastMap[index]-firstMap[index]+1)
		}
	}
	return ans
}
