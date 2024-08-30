package leetcode

// 执行耗时:117 ms,击败了43.90% 的Go用户
// 内存消耗:7.8 MB,击败了90.24% 的Go用户
func sumDigitDifferences(nums []int) (result int64) {
	var countArr = [10][10]int{}
	for i := range nums {
		var index = 0
		for nums[i] != 0 {
			countArr[index][nums[i]%10]++
			nums[i] /= 10
			index++
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			result += int64(countArr[i][j] * (len(nums) - countArr[i][j]))
		}
	}
	result /= 2
	return
}
