package leetcode

import "sort"

// 执行耗时:144 ms,击败了65.12% 的Go用户
// 内存消耗:17.1 MB,击败了30.23% 的Go用户
func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	var right = points[0][0] + w
	var result = 1
	for index := range points {
		if points[index][0] > right {
			result++
			right = points[index][0] + w
		}
	}
	return result
}
