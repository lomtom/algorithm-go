package leetcode

// 执行耗时:34 ms,击败了88.40% 的Go用户
// 内存消耗:7.87 MB,8.36% 的Go用户
func subarraySum(nums []int, k int) (result int) {
	sums := make(map[int]int)
	sums[0] = 1
	sum := 0
	for _, num := range nums {
		sum += num
		if v, ok := sums[sum-k]; ok {
			result += v
		}
		sums[sum]++
	}
	return
}
