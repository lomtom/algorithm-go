package leetcode

// 执行耗时:19 ms,击败了50.00% 的Go用户
// 内存消耗:6.6 MB,击败了34.38% 的Go用户
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	var nums1Map = make(map[int]int)
	var nums2Map = make(map[int]int)
	for _, num := range nums1 {
		nums1Map[num]++
	}
	for _, num := range nums2 {
		nums2Map[num]++
	}
	var result = make([]int, 2)
	for num, count := range nums1Map {
		if _, ok := nums2Map[num]; ok {
			result[0] += count
			result[1] += nums2Map[num]
		}
	}
	return result
}
